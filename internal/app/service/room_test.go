package service

import (
	"context"
	"github.com/mhrlife/tonference/internal/ent/room"
	"github.com/mhrlife/tonference/pkg/framework/apperror"
	"github.com/teris-io/shortid"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RoomTestSuite struct {
	suite.Suite

	*Service
	ctx context.Context
}

func TestRoomSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(RoomTestSuite))
}

func (s *RoomTestSuite) SetupSuite() {
	s.Service = NewTestingService(s.T())
	s.ctx = context.Background()
}

func (s *RoomTestSuite) TearDownSuite() {
	s.Service.Close()
}

func (s *RoomTestSuite) TestCreateRoom() {
	_, err := s.client.User.Create().SetDisplayName("User").SetGameID(shortid.MustGenerate()).Save(s.ctx)
	s.Require().NoError(err)

	dto := CreateRoomDto{CreatorID: 1}
	createdRoom, err := s.CreateRoom(s.ctx, dto)

	s.Require().NoError(err)
	s.Require().NotEmpty(createdRoom.ShortID)

	// Check if user is added to the room
	fetchedRoom, err := s.client.Room.Query().WithAdmins().Where(room.IDEQ(createdRoom.ID)).Only(s.ctx)
	s.Require().NoError(err)
	s.Require().Len(fetchedRoom.Edges.Admins, 1)
}

func (s *RoomTestSuite) TestConstraintMaxRoom() {
	_, err := s.client.User.Create().SetDisplayName("User").SetGameID(shortid.MustGenerate()).Save(s.ctx)
	s.Require().NoError(err)

	roomsToBeCreated := 3

	for i := 0; i < roomsToBeCreated; i++ {
		dto := CreateRoomDto{CreatorID: 1}
		_, err := s.CreateRoom(s.ctx, dto)
		s.Require().NoError(err)
	}

	dto := CreateRoomDto{CreatorID: 1}
	_, err = s.CreateRoom(s.ctx, dto)
	s.Require().Error(err)
	s.Require().True(apperror.IsConstraintError(err))
}

func (s *RoomTestSuite) TestRemoveRoom() {
	_, err := s.client.User.Create().SetDisplayName("User").SetGameID(shortid.MustGenerate()).Save(s.ctx)
	s.Require().NoError(err)

	dto := CreateRoomDto{CreatorID: 1}
	createdRoom, err := s.CreateRoom(s.ctx, dto)
	s.Require().NoError(err)

	// check if others can delete room
	_, err = s.client.User.Create().SetDisplayName("User").SetGameID(shortid.MustGenerate()).Save(s.ctx)
	s.Require().NoError(err)
	err = s.DeleteRoom(s.ctx, createdRoom.ShortID, 2)
	s.Require().Error(err)

	err = s.DeleteRoom(s.ctx, createdRoom.ShortID, 1)
	s.Require().NoError(err)

	deletedRoom, err := s.client.Room.Query().Where(room.IDEQ(createdRoom.ID)).Only(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal(deletedRoom.Status, room.StatusInactive)
}
