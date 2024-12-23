package service

import (
	"context"
	"github.com/livekit/protocol/auth"
	"github.com/mhrlife/tonference/internal/ent"
	"github.com/mhrlife/tonference/internal/ent/room"
	"github.com/mhrlife/tonference/internal/ent/user"
	"github.com/mhrlife/tonference/pkg/framework/apperror"
	"github.com/mhrlife/tonference/pkg/pt"
	"strconv"
	"time"
)

type CreateRoomDto struct {
	CreatorID int64
}

func (s *Service) CreateRoom(ctx context.Context, dto CreateRoomDto) (createdRoom *ent.Room, err error) {

	err = s.WithTx(ctx, func(tx *ent.Tx) error {

		hasMaxRoomExceeded, err := s.HasMaxRoomExceeded(tx, ctx, dto.CreatorID)
		if err != nil {
			return apperror.Wrap(err, "failed to check max room exceeded")
		}

		if hasMaxRoomExceeded {
			return apperror.NewConstraintError("Max room count exceeded, please delete some rooms")
		}

		createdRoom, err = tx.Room.Create().AddAdminIDs(dto.CreatorID).Save(ctx)
		if err != nil {
			return apperror.Wrap(err, "failed to create room")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdRoom, nil
}

func (s *Service) HasMaxRoomExceeded(tx *ent.Tx, ctx context.Context, userID int64) (bool, error) {
	roomCount, err := tx.Room.Query().Where(
		room.HasAdminsWith(user.IDEQ(userID)),
		room.StatusEQ(room.StatusActive),
	).
		Count(ctx)
	if err != nil {
		return false, apperror.Wrap(err, "failed to count rooms")
	}

	return roomCount >= 3, nil
}

func (s *Service) MyRooms(ctx context.Context, userID int64) ([]*ent.Room, error) {
	rooms, err := s.client.Room.Query().Where(
		room.HasAdminsWith(user.IDEQ(userID)),
		room.StatusEQ(room.StatusActive),
	).All(ctx)
	if err != nil {
		return nil, apperror.Wrap(err, "failed to get rooms")
	}

	return rooms, nil
}

func (s *Service) DeleteRoom(ctx context.Context, roomID string, userID int64) error {
	count, err := s.client.Room.Update().Where(
		room.HasAdminsWith(user.IDEQ(userID)),
		room.StatusEQ(room.StatusActive),
		room.ShortIDEQ(roomID),
	).SetStatus(room.StatusInactive).Save(ctx)
	if err != nil {
		return apperror.Wrap(err, "failed to delete room")
	}

	if count == 0 {
		return apperror.NewNotFoundError("you are not allowed to delete this room")
	}

	return nil
}

func (s *Service) GenerateRoomToken(ctx context.Context, roomID string, user *ent.User) (string, error) {
	requestedRoom, err := s.client.Room.Query().WithAdmins().Where(
		room.StatusEQ(room.StatusActive),
		room.ShortIDEQ(roomID),
	).Only(ctx)
	if err != nil {
		return "", apperror.Wrap(err, "failed to get room")
	}

	isAdmin := false
	for _, admin := range requestedRoom.Edges.Admins {
		if admin.ID != user.ID {
			continue
		}

		isAdmin = true
		break
	}

	token, err := s.liveKitRoomService.CreateToken().SetIdentity(strconv.FormatInt(user.ID, 10)).SetVideoGrant(&auth.VideoGrant{
		Room:         roomID,
		RoomJoin:     true,
		CanPublish:   pt.Value(true),
		CanSubscribe: pt.Value(true),
	}).SetAttributes(map[string]string{
		"room_id":      roomID,
		"game_id":      user.GameID,
		"is_admin":     strconv.FormatBool(isAdmin),
		"display_name": user.DisplayName,
	}).SetValidFor(time.Hour * 24).ToJWT()

	if err != nil {
		return "", apperror.Wrap(err, "failed to generate token")
	}

	return token, nil
}
