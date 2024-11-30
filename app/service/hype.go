// app/service/hype.go
package service

import (
	"context"
	"time"

	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
	"nevissGo/ent"
	"nevissGo/framework"
)

type Hype struct {
	app            *framework.App
	client         *ent.Client
	defaultMaxHype int
}

func NewHype(app *framework.App) *Hype {
	return &Hype{
		app:            app,
		client:         app.Client(),
		defaultMaxHype: 100,
	}
}

func (h *Hype) UseHypeTX(ctx context.Context, tx *ent.Tx, userID int64, amount int) error {
	user, err := tx.User.Get(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			logrus.WithField("user_id", userID).Error("User not found")
			return eris.Errorf("user with ID %d not found", userID)
		}
		logrus.WithError(err).WithField("user_id", userID).Error("Failed to get user")
		return eris.Wrapf(err, "failed to get user with ID %d", userID)
	}

	hype, err := user.QueryHype().Only(ctx)
	if ent.IsNotFound(err) {
		hype, err = tx.Hype.Create().
			SetUser(user).
			SetAmountRemaining(h.defaultMaxHype).
			SetMaxHype(h.defaultMaxHype).
			SetHypePerMinute(2).
			SetLastUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			logrus.WithError(err).WithField("user_id", userID).Error("Failed to create hype for user")
			return eris.Wrap(err, "failed to create hype for user")
		}
	} else if err != nil {
		logrus.WithError(err).WithField("user_id", userID).Error("Failed to query hype for user")
		return eris.Wrap(err, "failed to query hype for user")
	} else {
		timeSinceUpdate := time.Since(hype.LastUpdatedAt)
		minutesPassed := int(timeSinceUpdate.Minutes())
		if minutesPassed > 0 {
			replenished := minutesPassed * hype.HypePerMinute
			newAmount := hype.AmountRemaining + replenished
			if newAmount > hype.MaxHype {
				newAmount = hype.MaxHype
			}
			hype.AmountRemaining = newAmount
			hype.LastUpdatedAt = hype.LastUpdatedAt.Add(time.Duration(minutesPassed) * time.Minute)
		}
	}

	if hype.AmountRemaining < amount {
		logrus.WithFields(logrus.Fields{
			"user_id":          userID,
			"amount_required":  amount,
			"amount_remaining": hype.AmountRemaining,
		}).Warn("Not enough hype remaining for user")
		return eris.New("not enough hype remaining")
	}

	hype.AmountRemaining -= amount
	hype.LastUpdatedAt = time.Now()

	_, err = tx.Hype.UpdateOne(hype).
		SetAmountRemaining(hype.AmountRemaining).
		SetLastUpdatedAt(hype.LastUpdatedAt).
		Save(ctx)
	if err != nil {
		logrus.WithError(err).WithField("user_id", userID).Error("Failed to update hype")
		return eris.Wrap(err, "failed to update hype")
	}

	logrus.WithFields(logrus.Fields{
		"user_id":          userID,
		"amount_used":      amount,
		"amount_remaining": hype.AmountRemaining,
	}).Info("Hype used successfully")

	return nil
}
