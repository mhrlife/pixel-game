package service

import (
	"context"
	"nevissGo/ent/pixel"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"nevissGo/ent"
	"nevissGo/framework"
)

type PixelsSuite struct {
	suite.Suite
	app      *framework.TestingApp
	service  *Pixels
	ctx      context.Context
	cooldown time.Duration
	user     *ent.User
	bridge   TestingBridge
}

func TestPixelsSuite(t *testing.T) {
	suite.Run(t, new(PixelsSuite))
}

func (s *PixelsSuite) SetupTest() {
	s.app = framework.NewTestingApp(s.T())
	s.bridge = TestBridge(s.T())
	s.cooldown = 2 * time.Second
	s.service = NewPixels(s.app.App, s.bridge.Bridge, s.cooldown, 10, 10, 1)
	s.ctx = context.Background()

	var err error
	s.user, err = s.app.Client().User.Create().
		SetDisplayName("TestUser").
		SetGameID("game123").
		Save(s.ctx)
	s.NoError(err)
}

func (s *PixelsSuite) TestUpdateColorCreatePixel() {
	validPixelID := 5
	newColor := "green"

	s.bridge.Hype.On("UseHypeTX", mock.Anything, mock.Anything, s.user.ID, 1).Return(nil)
	defer s.bridge.Hype.AssertExpectations(s.T())

	err := s.service.UpdateColor(s.ctx, validPixelID, newColor, s.user.ID)

	s.NoError(err)

	createdPixel, err := s.app.Client().Pixel.
		Query().
		Where(pixel.IDEQ(validPixelID)).
		WithUser().
		Only(s.ctx)
	s.NoError(err)
	s.Equal(newColor, createdPixel.Color)
	s.Equal(s.user.ID, createdPixel.Edges.User.ID)
}

func (s *PixelsSuite) TestUpdateColorUpdateExistingPixel() {
	pixelID := 3
	existingColor := "red"
	newColor := "blue"

	err := s.app.TX(s.ctx, func(tx *ent.Tx) error {
		_, err := tx.Pixel.Create().
			SetID(pixelID).
			SetColor(existingColor).
			SetUpdatedAt(time.Now().Add(-3 * time.Second)).
			SetUserID(s.user.ID).
			Save(s.ctx)
		return err
	})
	s.NoError(err)

	s.bridge.Hype.On("UseHypeTX", mock.Anything, mock.Anything, s.user.ID, 1).Return(nil)
	defer s.bridge.Hype.AssertExpectations(s.T())

	err = s.service.UpdateColor(s.ctx, pixelID, newColor, s.user.ID)

	s.NoError(err)

	updatedPixel, err := s.app.Client().Pixel.
		Query().
		Where(pixel.IDEQ(pixelID)).
		WithUser().
		Only(s.ctx)
	s.NoError(err)
	s.Equal(newColor, updatedPixel.Color)
	s.Equal(s.user.ID, updatedPixel.Edges.User.ID)
}

func (s *PixelsSuite) TestUpdateColorCooldownNotExpired() {
	pixelID := 2
	existingColor := "yellow"

	err := s.app.TX(s.ctx, func(tx *ent.Tx) error {
		_, err := tx.Pixel.Create().
			SetID(pixelID).
			SetColor(existingColor).
			SetUpdatedAt(time.Now()).
			SetUserID(s.user.ID).
			Save(s.ctx)
		return err
	})
	s.NoError(err)

	err = s.service.UpdateColor(s.ctx, pixelID, "purple", s.user.ID)

	s.Error(err)
	s.Equal(400, framework.ExtErrorCode(err))
	s.Equal("Pixel can only be updated every "+s.cooldown.String(), framework.ExtErrorMessage(err))
}

func (s *PixelsSuite) TestUpdateColorInvalidPixelID() {
	invalidPixelID := 100
	newColor := "black"

	err := s.service.UpdateColor(s.ctx, invalidPixelID, newColor, s.user.ID)

	s.Error(err)
	s.Equal(400, framework.ExtErrorCode(err))
	s.Equal("Pixel ID is out of bounds", framework.ExtErrorMessage(err))
}

func (s *PixelsSuite) TestUpdateColorUseHypeFailure() {
	pixelID := 4
	newColor := "orange"

	s.bridge.Hype.On("UseHypeTX", mock.Anything, mock.Anything, s.user.ID, 1).Return(framework.NewInternalError("hype usage failed"))
	defer s.bridge.Hype.AssertExpectations(s.T())

	err := s.service.UpdateColor(s.ctx, pixelID, newColor, s.user.ID)

	s.Error(err)
	s.Equal(500, framework.ExtErrorCode(err))
	s.Equal("hype usage failed", framework.ExtErrorMessage(err))
}

func (s *PixelsSuite) TestGetBoard() {
	err := s.app.TX(s.ctx, func(tx *ent.Tx) error {
		_, err := tx.Pixel.Create().
			SetID(1).
			SetColor("red").
			SetUpdatedAt(time.Now()).
			SetUserID(s.user.ID).
			Save(s.ctx)
		return err
	})
	s.NoError(err)

	board, err := s.service.GetBoard(s.ctx)
	s.NoError(err)
	s.Equal(10, board.Width)
	s.Equal(10, board.Height)
	s.Len(board.Pixels, 100)
	s.Equal("red", board.Pixels[1].Color)
	s.Equal("white", board.Pixels[0].Color)
}
