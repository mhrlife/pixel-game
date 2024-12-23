package serializer

import "github.com/mhrlife/tonference/internal/ent"

type RoomSerializer struct {
	ID              string `json:"id"`
	MaxParticipants int8   `json:"max_participants"`

	PermissionVideo string `json:"permission_video"`
	PermissionAudio string `json:"permission_audio"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func NewRoomSerializer(room *ent.Room) *RoomSerializer {
	return &RoomSerializer{
		ID:              room.ShortID,
		MaxParticipants: room.MaxParticipants,
		PermissionVideo: room.PermissionVideo.String(),
		PermissionAudio: room.PermissionAudio.String(),
		CreatedAt:       room.CreatedAt.Unix(),
		UpdatedAt:       room.UpdatedAt.Unix(),
	}
}
