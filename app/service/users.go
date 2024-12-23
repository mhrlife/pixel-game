package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
	"nevissGo/ent"
	"nevissGo/framework"
)

type Users struct {
	app *framework.App
}

func NewUsers(app *framework.App) *Users {
	return &Users{
		app: app,
	}
}

func (s *Users) GetOrRegister(ctx context.Context, user *ent.User) error {
	return s.app.TX(ctx, func(tx *ent.Tx) error {
		existingUser, err := tx.User.Get(ctx, user.ID)
		if ent.IsNotFound(err) {
			_, err = tx.User.Create().
				SetID(user.ID).
				SetDisplayName(user.DisplayName).
				SetGameID(shortid.MustGenerate()).
				Save(ctx)
			if err != nil {
				logrus.WithError(err).Error("Failed to create user")
				return framework.NewInternalError("Failed to create user")
			}
			return nil
		}

		if err != nil {
			logrus.WithError(err).Error("Failed to get user")
			return framework.NewInternalError("Failed to get user")
		}

		user = existingUser
		return nil
	})
}
func (s *Users) Get(ctx context.Context, userID int64) (*ent.User, error) {
	return s.app.Client().User.Get(ctx, userID)
}
