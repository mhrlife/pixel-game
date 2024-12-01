package service

import (
	"context"
	"time"

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
		defaultMaxHype: 10,
	}
}

func (h *Hype) UseHypeTX(ctx context.Context, tx *ent.Tx, userID int64, amount int) error {
	hype, err := h.fetchOrCreateHype(ctx, tx.Client(), userID)
	if err != nil {
		return err
	}

	err = h.updateHypeAmount(ctx, tx.Client(), hype)
	if err != nil {
		return err
	}

	if hype.AmountRemaining < amount {
		return framework.NewValidationError("not enough hype remaining")
	}

	hype.AmountRemaining -= amount
	hype.LastUpdatedAt = time.Now()

	_, err = tx.Hype.UpdateOne(hype).
		SetAmountRemaining(hype.AmountRemaining).
		SetLastUpdatedAt(hype.LastUpdatedAt).
		Save(ctx)
	if err != nil {
		return framework.NewInternalError("Failed to update hype")
	}

	return nil
}

func (h *Hype) GetHype(ctx context.Context, userID int64) (*ent.Hype, error) {
	hype, err := h.fetchOrCreateHype(ctx, h.client, userID)
	if err != nil {
		return nil, err
	}

	err = h.updateHypeAmount(ctx, h.client, hype)
	if err != nil {
		return nil, err
	}

	return hype, nil
}

func (h *Hype) fetchOrCreateHype(ctx context.Context, client *ent.Client, userID int64) (*ent.Hype, error) {
	user, err := client.User.Get(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, framework.NewValidationError("User not found")
		}
		return nil, framework.NewInternalError("Failed to get user")
	}

	hype, err := user.QueryHype().Only(ctx)
	if ent.IsNotFound(err) {
		hype, err = client.Hype.Create().
			SetUser(user).
			SetAmountRemaining(h.defaultMaxHype).
			SetMaxHype(h.defaultMaxHype).
			SetHypePerMinute(2).
			SetLastUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, framework.NewInternalError("Failed to create hype for user")
		}
	} else if err != nil {
		return nil, framework.NewInternalError("Failed to query hype for user")
	}

	return hype, nil
}

func (h *Hype) updateHypeAmount(ctx context.Context, client *ent.Client, hype *ent.Hype) error {
	timeSinceUpdate := time.Since(hype.LastUpdatedAt)
	hypePerSecond := float64(hype.HypePerMinute) / 60.0
	secondsPassed := timeSinceUpdate.Seconds()
	replenished := int(secondsPassed * hypePerSecond)
	if replenished > 0 {
		newAmount := hype.AmountRemaining + replenished
		if newAmount > hype.MaxHype {
			newAmount = hype.MaxHype
		}
		hype.AmountRemaining = newAmount
		hype.LastUpdatedAt = hype.LastUpdatedAt.Add(time.Duration(float64(time.Second) * float64(replenished) / hypePerSecond))
		_, err := client.Hype.UpdateOne(hype).
			SetAmountRemaining(hype.AmountRemaining).
			SetLastUpdatedAt(hype.LastUpdatedAt).
			Save(ctx)
		if err != nil {
			return framework.NewInternalError("Failed to update hype amount")
		}
	}
	return nil
}
