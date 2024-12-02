package serializer

import (
	"nevissGo/app/service"
	"nevissGo/ent"
	"time"
)

type PixelSerializer struct {
	ID        int    `json:"id"`
	Color     string `json:"color"`
	UpdatedAt string `json:"updated_at"`
}

type PixelWithUserSerializer struct {
	ID        int    `json:"id"`
	Color     string `json:"color"`
	User      *User  `json:"user,omitempty"`
	UpdatedAt int64  `json:"updated_at"`
}

type BoardSerializer struct {
	Pixels    []*PixelWithUserSerializer `json:"pixels"`
	Width     int                        `json:"width"`
	Height    int                        `json:"height"`
	UpdatedAt int64                      `json:"updated_at"`
}

func NewPixel(pixel *ent.Pixel) *PixelSerializer {
	return &PixelSerializer{
		ID:    pixel.ID,
		Color: pixel.Color,
	}
}

func NewPixelWithUser(pixel *ent.Pixel) *PixelWithUserSerializer {
	var user *User
	if pixel.Edges.User != nil {
		u := NewUser(pixel.Edges.User)
		user = &u
	}

	return &PixelWithUserSerializer{
		ID:        pixel.ID,
		Color:     pixel.Color,
		User:      user,
		UpdatedAt: pixel.UpdatedAt.Unix(),
	}
}

func NewBoard(board *service.Board) *BoardSerializer {
	pixels := make([]*PixelWithUserSerializer, board.Width*board.Height)
	for _, pixel := range board.Pixels {
		if pixel != nil {
			pixels[pixel.ID] = NewPixelWithUser(pixel)
		}
	}

	minUpdatedAt := time.Now().Unix()

	for i := 0; i < len(pixels); i++ {
		if pixels[i] == nil {
			pixels[i] = &PixelWithUserSerializer{
				ID:        i,
				Color:     "white",
				UpdatedAt: time.Now().Unix(),
			}
		}

		if pixels[i].UpdatedAt < 0 {
			pixels[i].UpdatedAt = time.Now().Unix()
		}

		if pixels[i].UpdatedAt < minUpdatedAt {
			minUpdatedAt = pixels[i].UpdatedAt
		}
	}
	return &BoardSerializer{
		Pixels:    pixels,
		Width:     board.Width,
		Height:    board.Height,
		UpdatedAt: minUpdatedAt,
	}
}

type UpdatedBoardSerializer struct {
	Board *BoardSerializer `json:"board"`
	User  User             `json:"user"`
}

func NewBoardUpdatedSerializer(board *service.Board, user *ent.User) *UpdatedBoardSerializer {
	return &UpdatedBoardSerializer{
		Board: NewBoard(board),
		User:  NewUser(user),
	}
}
