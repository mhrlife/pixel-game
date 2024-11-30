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
			user, err = tx.User.Create().
				SetID(user.ID).
				SetDisplayName(user.DisplayName).
				SetGameID(shortid.MustGenerate()).
				Save(ctx)
			return err
		}

		if err != nil {
			logrus.WithError(err).Error("failed to get service")
			return err
		}

		user = existingUser
		return nil
	})
}
