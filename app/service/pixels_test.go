// app/service/pixels_test.go
package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"nevissGo/ent"
	"nevissGo/framework"
)

// PixelsSuite defines the test suite for the Pixels service.
type PixelsSuite struct {
	suite.Suite
	app      *framework.TestingApp
	service  *Pixels
	ctx      context.Context
	cooldown time.Duration
}

// TestPixelsSuite runs the Pixels test suite.
func TestPixelsSuite(t *testing.T) {
	suite.Run(t, new(PixelsSuite))
}

func (s *PixelsSuite) SetupTest() {
	s.app = framework.NewTestingApp(s.T())
	s.cooldown = 2 * time.Second // Example cooldown duration
	s.service = NewPixels(s.app.App, s.cooldown, 3, 3)
	s.ctx = context.Background()
}

func (s *PixelsSuite) TestUpdateColorCreatePixel() {
	// Arrange: Define a pixelID within the valid range but not present in the database
	validPixelID := 5 // Assuming width*height > 5
	newColor := "green"

	// Act: Attempt to update the color of the non-existent pixel
	err := s.service.UpdateColor(s.ctx, validPixelID, newColor)

	// Assert: Ensure no error occurred and the pixel was created with the new color
	s.NoError(err)

	createdPixel, err := s.app.Client().Pixel.Get(s.ctx, validPixelID)
	s.NoError(err)
	s.Equal(newColor, createdPixel.Color)
}

func (s *PixelsSuite) TestUpdateColorOutOfBounds() {
	// Arrange: Define a pixelID outside the valid range
	invalidPixelID := s.service.width * s.service.height // One more than the maximum valid ID

	// Act: Attempt to update the color of the out-of-bounds pixel
	err := s.service.UpdateColor(s.ctx, invalidPixelID, "blue")

	// Assert: Ensure an error is returned indicating the pixelID is out of bounds
	s.Error(err)
	s.Contains(err.Error(), "Pixel ID")
	s.Contains(err.Error(), "out of bounds")
}

func (s *PixelsSuite) TestUpdateColorBoundaryPixel() {
	// Arrange: Define the maximum valid pixelID
	maxValidPixelID := s.service.width*s.service.height - 1
	newColor := "purple"

	// Act: Attempt to update the color of the boundary pixel
	err := s.service.UpdateColor(s.ctx, maxValidPixelID, newColor)

	// Assert: Ensure no error occurred and the pixel was updated/created correctly
	s.NoError(err)

	updatedPixel, err := s.app.Client().Pixel.Get(s.ctx, maxValidPixelID)
	s.NoError(err)
	s.Equal(newColor, updatedPixel.Color)
}

// TestUpdateColorPixelNotFound tests updating a color for a non-existent pixel.
func (s *PixelsSuite) TestUpdateColorPixelNotFound() {
	// Arrange: Ensure the pixel does not exist
	nonExistentPixelID := 999

	// Act: Attempt to update the pixel's color
	err := s.service.UpdateColor(s.ctx, nonExistentPixelID, "green")

	// Assert: Ensure an error is returned indicating the pixel was not found
	s.Error(err)
	s.Contains(err.Error(), "out of bounds")
}

// TestUpdateColorCooldownNotPassed tests updating a pixel's color before the cooldown period has passed.
func (s *PixelsSuite) TestUpdateColorCooldownNotPassed() {
	// Arrange: Create a pixel recently updated
	pixel := &ent.Pixel{
		ID:        2,
		Color:     "yellow",
		UpdatedAt: time.Now().Add(-s.cooldown / 2),
	}
	_, err := s.app.Client().Pixel.Create().
		SetID(pixel.ID).
		SetColor(pixel.Color).
		SetUpdatedAt(pixel.UpdatedAt).
		Save(s.ctx)
	s.NoError(err)

	// Act: Attempt to update the pixel's color
	err = s.service.UpdateColor(s.ctx, pixel.ID, "purple")

	// Assert: Ensure an error is returned indicating the cooldown period has not passed
	s.Error(err)
	s.Contains(err.Error(), "Pixel can only be updated")
}

// TestUpdateColorConcurrentUpdates tests handling of concurrent updates to the same pixel.
func (s *PixelsSuite) TestUpdateColorConcurrentUpdates() {
	// Arrange: Create a pixel in the database
	pixel := &ent.Pixel{
		ID:        4,
		Color:     "orange",
		UpdatedAt: time.Now().Add(-s.cooldown * 2),
	}
	_, err := s.app.Client().Pixel.Create().
		SetID(pixel.ID).
		SetColor(pixel.Color).
		SetUpdatedAt(pixel.UpdatedAt).
		Save(s.ctx)
	s.NoError(err)

	// get created pixel
	pix, err := s.app.Client().Pixel.Get(s.ctx, pixel.ID)
	s.NoError(err)
	s.Equal(pixel.Color, pix.Color)

	// Act: Perform two concurrent updates
	newColor1 := "pink"
	newColor2 := "cyan"

	errCh := make(chan error, 2)
	go func() {
		errCh <- s.service.UpdateColor(s.ctx, pixel.ID, newColor1)
	}()
	go func() {
		errCh <- s.service.UpdateColor(s.ctx, pixel.ID, newColor2)
	}()

	// Assert: Only one update should succeed
	var firstErr, secondErr error
	firstErr = <-errCh
	secondErr = <-errCh

	successCount := 0
	if firstErr == nil {
		successCount++
	}
	if secondErr == nil {
		successCount++
	}
	s.Equal(1, successCount, "only one update should succeed due to cooldown")

	// Verify the final color is one of the new colors
	updatedPixel, err := s.app.Client().Pixel.Get(s.ctx, pixel.ID)
	s.NoError(err)
	s.Contains([]string{newColor1, newColor2}, updatedPixel.Color)
}

// TestGetBoard tests retrieving the board of pixels.
func (s *PixelsSuite) TestGetBoard() {
	// Arrange: Create a board with 3x3 pixels
	board := &Board{
		Width:  3,
		Height: 3,
	}

	for i := 0; i < board.Width; i++ {
		_, err := s.app.Client().Pixel.Create().
			SetID(i).
			SetColor("white").
			SetUpdatedAt(time.Now()).
			Save(s.ctx)
		s.NoError(err)
	}

	// Act: Retrieve the board
	retrievedBoard, err := s.service.GetBoard(s.ctx)

	// Assert: Ensure no error occurred and the board has the correct number of pixels
	s.NoError(err)
	s.Len(retrievedBoard.Pixels, board.Width*board.Height)
	s.Equal(board.Width, retrievedBoard.Width)
	s.Equal(board.Height, retrievedBoard.Height)

	// Assert: Ensure all pixels are present in the board
	for i := range retrievedBoard.Pixels {
		s.NotNil(retrievedBoard.Pixels[i])
	}
}
