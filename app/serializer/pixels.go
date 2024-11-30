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
	UpdatedAt string `json:"updated_at"`
}

type BoardSerializer struct {
	Pixels []*PixelWithUserSerializer `json:"pixels"`
	Width  int                        `json:"width"`
	Height int                        `json:"height"`
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

	updatedAt := ""
	if !pixel.UpdatedAt.IsZero() {
		updatedAt = pixel.UpdatedAt.Format(time.RFC3339)
	}

	return &PixelWithUserSerializer{
		ID:        pixel.ID,
		Color:     pixel.Color,
		User:      user,
		UpdatedAt: updatedAt,
	}
}

func NewBoard(board *service.Board) *BoardSerializer {
	pixels := make([]*PixelWithUserSerializer, board.Width*board.Height)
	for _, pixel := range board.Pixels {
		if pixel != nil {
			pixels[pixel.ID] = NewPixelWithUser(pixel)
		}
	}
	for i := 0; i < len(pixels); i++ {
		if pixels[i] == nil {
			pixels[i] = &PixelWithUserSerializer{
				ID:    i,
				Color: "white",
			}
		}
	}
	return &BoardSerializer{
		Pixels: pixels,
		Width:  board.Width,
		Height: board.Height,
	}
}
