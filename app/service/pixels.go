// app/service/pixels.go
package service

import (
	"context"
	"github.com/rotisserie/eris"
	"time"

	"github.com/sirupsen/logrus"
	"nevissGo/ent"
	"nevissGo/framework"
)

// Pixels handles operations related to Pixel entities.
type Pixels struct {
	app      *framework.App
	cooldown time.Duration

	width, height int
}

func NewPixels(app *framework.App, cooldown time.Duration, width, height int) *Pixels {
	return &Pixels{
		app:      app,
		cooldown: cooldown,
		width:    width,
		height:   height,
	}
}

// UpdateColor updates the color of a pixel. If the pixel does not exist and the pixelID
// is within the valid range, it creates the pixel with the specified color.
// Returns an error if the pixelID is out of bounds or if any database operation fails.
func (s *Pixels) UpdateColor(ctx context.Context, pixelID int, newColor string) error {
	return s.app.TX(ctx, func(tx *ent.Tx) error {
		existingPixel, err := tx.Pixel.Get(ctx, pixelID)
		if err != nil {
			if ent.IsNotFound(err) {
				// Check if the pixelID is within the valid range
				if pixelID < 0 || pixelID >= s.width*s.height {
					logrus.WithFields(logrus.Fields{
						"pixel_id": pixelID,
						"width":    s.width,
						"height":   s.height,
					}).Error("Pixel ID is out of bounds")
					return eris.Errorf("Pixel ID %d is out of bounds (0 to %d)", pixelID, s.width*s.height-1)
				}

				// Create the pixel since it doesn't exist and pixelID is valid
				_, err := tx.Pixel.Create().
					SetID(pixelID).
					SetColor(newColor).
					Save(ctx)
				if err != nil {
					logrus.WithError(err).WithField("pixel_id", pixelID).Error("Failed to create pixel")
					return eris.Wrapf(err, "failed to create pixel with ID %d", pixelID)
				}

				logrus.WithFields(logrus.Fields{
					"pixel_id":  pixelID,
					"new_color": newColor,
				}).Info("Pixel created successfully")
				return nil
			}

			// Log and return other errors encountered while retrieving the pixel
			logrus.WithError(err).WithField("pixel_id", pixelID).Error("Failed to retrieve pixel")
			return eris.Wrapf(err, "failed to retrieve pixel with ID %d", pixelID)
		}

		// Calculate time since the last update
		timeSinceUpdate := time.Since(existingPixel.UpdatedAt)
		if timeSinceUpdate < s.cooldown {
			logrus.WithFields(logrus.Fields{
				"pixel_id":          pixelID,
				"time_since_update": timeSinceUpdate.Seconds(),
				"cooldown_secs":     s.cooldown.Seconds(),
			}).Warn("Attempt to update pixel too soon")
			return eris.Errorf("Pixel can only be updated every %s", s.cooldown)
		}

		// Update the pixel's color
		_, err = tx.Pixel.UpdateOne(existingPixel).
			SetColor(newColor).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			logrus.WithError(err).WithField("pixel_id", pixelID).Error("Failed to update pixel color")
			return eris.Wrapf(err, "failed to update color for pixel with ID %d", pixelID)
		}

		logrus.WithFields(logrus.Fields{
			"pixel_id":  pixelID,
			"new_color": newColor,
		}).Info("Pixel color updated successfully")

		return nil
	})
}
	
type Board struct {
	Pixels []*ent.Pixel // Pixels index is *ent.Pixel's ID
	Width  int
	Height int
}

func (s *Pixels) GetBoard(ctx context.Context) (*Board, error) {
	board := &Board{
		Width:  s.width,
		Height: s.height,
	}

	pixels, err := s.app.Client().Pixel.Query().All(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "failed to retrieve pixels")
	}

	board.Pixels = make([]*ent.Pixel, s.width*s.height)
	for _, pixel := range pixels {
		board.Pixels[pixel.ID] = pixel
	}

	// Fill in any missing pixels
	for i := range board.Pixels {
		if board.Pixels[i] == nil {
			board.Pixels[i] = &ent.Pixel{ID: i, Color: "white"}
		}
	}

	return board, nil
}
