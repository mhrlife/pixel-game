package service

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"nevissGo/ent"
	"nevissGo/framework"
)

type Pixels struct {
	app           *framework.App
	bridge        Bridge
	cooldown      time.Duration
	width, height int
	drawHypeCost  int
}

func NewPixels(
	app *framework.App,
	bridge Bridge,

	cooldown time.Duration,
	width, height int,
	drawHypeCost int,
) *Pixels {
	return &Pixels{
		app:    app,
		bridge: bridge,

		cooldown:     cooldown,
		width:        width,
		height:       height,
		drawHypeCost: drawHypeCost,
	}
}

func (s *Pixels) UpdateColor(ctx context.Context, pixelID int, newColor string, userID int64) error {
	if pixelID < 0 || pixelID >= s.width*s.height {
		logrus.WithFields(logrus.Fields{
			"pixel_id": pixelID,
			"width":    s.width,
			"height":   s.height,
		}).Error("Pixel ID is out of bounds")
		return framework.NewValidationError("Pixel ID is out of bounds")
	}

	return s.app.TX(ctx, func(tx *ent.Tx) error {
		pixel, err := s.getPixel(tx, ctx, pixelID)
		if ent.IsNotFound(err) {
			if err := s.bridge.Hype.UseHypeTX(ctx, tx, userID, s.drawHypeCost); err != nil {
				return err
			}
			return s.createPixel(tx, ctx, pixelID, newColor, userID)
		}
		if err != nil {
			logrus.WithError(err).WithField("pixel_id", pixelID).Error("Failed to retrieve pixel")
			return framework.NewInternalError("Failed to retrieve pixel")
		}
		if err := s.ensureCooldown(pixel); err != nil {
			return err
		}
		if err := s.bridge.Hype.UseHypeTX(ctx, tx, userID, s.drawHypeCost); err != nil {
			return err
		}
		return s.updateExistingPixel(tx, ctx, pixel, newColor, userID)
	})
}

func (s *Pixels) getPixel(tx *ent.Tx, ctx context.Context, pixelID int) (*ent.Pixel, error) {
	return tx.Pixel.Get(ctx, pixelID)
}

func (s *Pixels) createPixel(tx *ent.Tx, ctx context.Context, pixelID int, newColor string, userID int64) error {
	if pixelID < 0 || pixelID >= s.width*s.height {
		logrus.WithFields(logrus.Fields{
			"pixel_id": pixelID,
			"width":    s.width,
			"height":   s.height,
		}).Error("Pixel ID is out of bounds")
		return framework.NewValidationError("Pixel ID is out of bounds")
	}
	_, err := tx.Pixel.Create().
		SetID(pixelID).
		SetColor(newColor).
		SetUpdatedAt(time.Now()).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		logrus.WithError(err).WithField("pixel_id", pixelID).Error("Failed to create pixel")
		return framework.NewInternalError("Failed to create pixel")
	}
	logrus.WithFields(logrus.Fields{
		"pixel_id":  pixelID,
		"new_color": newColor,
		"user_id":   userID,
	}).Info("Pixel created and assigned to user successfully")
	return nil
}

func (s *Pixels) ensureCooldown(pixel *ent.Pixel) error {
	timeSinceUpdate := time.Since(pixel.UpdatedAt)
	if timeSinceUpdate < s.cooldown {
		logrus.WithFields(logrus.Fields{
			"pixel_id":          pixel.ID,
			"time_since_update": timeSinceUpdate.Seconds(),
			"cooldown_secs":     s.cooldown.Seconds(),
		}).Warn("Attempt to update pixel too soon")
		return framework.NewValidationError("Pixel can only be updated every " + s.cooldown.String())
	}
	return nil
}

func (s *Pixels) updateExistingPixel(tx *ent.Tx, ctx context.Context, pixel *ent.Pixel, newColor string, userID int64) error {
	_, err := tx.Pixel.UpdateOne(pixel).
		SetColor(newColor).
		SetUpdatedAt(time.Now()).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		logrus.WithError(err).WithField("pixel_id", pixel.ID).Error("Failed to update pixel color")
		return framework.NewInternalError("Failed to update pixel color").WithFields(logrus.Fields{
			"pixel_id": pixel.ID,
		})
	}
	logrus.WithFields(logrus.Fields{
		"pixel_id":  pixel.ID,
		"new_color": newColor,
		"user_id":   userID,
	}).Info("Pixel color updated and reassigned to user successfully")
	return nil
}

type Board struct {
	Pixels []*ent.Pixel
	Width  int
	Height int
}

func (s *Pixels) GetBoard(ctx context.Context) (*Board, error) {
	board := &Board{
		Width:  s.width,
		Height: s.height,
	}

	pixels, err := s.app.Client().Pixel.Query().WithUser().All(ctx)
	if err != nil {
		return nil, framework.NewInternalError("Failed to retrieve pixels")
	}

	board.Pixels = make([]*ent.Pixel, s.width*s.height)
	for _, pixel := range pixels {
		board.Pixels[pixel.ID] = pixel
	}

	for i := range board.Pixels {
		if board.Pixels[i] == nil {
			board.Pixels[i] = &ent.Pixel{ID: i, Color: "white"}
		}
	}

	return board, nil
}
