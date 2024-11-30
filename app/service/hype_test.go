// app/service/hype_test.go
package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"nevissGo/ent"
	hype2 "nevissGo/ent/hype"
	"nevissGo/ent/user"
	"nevissGo/framework"
	"testing"
	"time"
)

type HypeSuite struct {
	suite.Suite
	bridge TestingBridge

	app     *framework.TestingApp
	service *Hype
	ctx     context.Context
	user    *ent.User
}

func TestHypeSuite(t *testing.T) {
	suite.Run(t, new(HypeSuite))
}

func (s *HypeSuite) SetupTest() {
	s.app = framework.NewTestingApp(s.T())
	s.bridge = TestBridge(s.T())
	s.service = NewHype(s.app.App)
	s.ctx = context.Background()

	var err error
	s.user, err = s.app.Client().User.Create().
		SetDisplayName("TestUser").
		SetGameID("game123").
		Save(s.ctx)
	s.NoError(err)
}

func (s *HypeSuite) TestUseHype_Success() {
	err := s.app.TX(s.ctx, func(tx *ent.Tx) error {
		_, err := tx.Hype.Create().
			SetUser(s.user).
			SetAmountRemaining(50).
			SetMaxHype(100).
			SetHypePerMinute(2).
			SetLastUpdatedAt(time.Now()).
			Save(s.ctx)
		return err
	})
	s.NoError(err)

	err = s.app.TX(s.ctx, func(tx *ent.Tx) error {
		return s.service.UseHypeTX(s.ctx, tx, s.user.ID, 30)
	})
	s.NoError(err)

	hype, err := s.app.Client().Hype.
		Query().
		Where(hype2.HasUserWith(user.IDEQ(s.user.ID))).
		Only(s.ctx)
	s.NoError(err)
	s.Equal(20, hype.AmountRemaining)
}

func (s *HypeSuite) TestUseHype_CreateHype() {
	count, err := s.app.Client().Hype.
		Query().
		Where(hype2.HasUserWith(user.IDEQ(s.user.ID))).
		Count(s.ctx)
	s.NoError(err)
	s.Equal(0, count)

	err = s.app.TX(s.ctx, func(tx *ent.Tx) error {
		return s.service.UseHypeTX(s.ctx, tx, s.user.ID, 40)
	})
	s.NoError(err)

	hype, err := s.app.Client().Hype.
		Query().
		Where(hype2.HasUserWith(user.IDEQ(s.user.ID))).
		Only(s.ctx)
	s.NoError(err)
	s.Equal(60, hype.AmountRemaining)
	s.Equal(100, hype.MaxHype)
}

func (s *HypeSuite) TestUseHype_UserNotFound() {
	nonExistentUserID := int64(9999)

	err := s.app.TX(s.ctx, func(tx *ent.Tx) error {
		return s.service.UseHypeTX(s.ctx, tx, nonExistentUserID, 10)
	})
	s.Error(err)
	s.Contains(err.Error(), "user with ID")
}

func (s *HypeSuite) TestUseHype_NotEnoughHype() {
	err := s.app.TX(s.ctx, func(tx *ent.Tx) error {
		_, err := tx.Hype.Create().
			SetUser(s.user).
			SetAmountRemaining(20).
			SetMaxHype(100).
			SetHypePerMinute(2).
			SetLastUpdatedAt(time.Now()).
			Save(s.ctx)
		return err
	})
	s.NoError(err)

	err = s.app.TX(s.ctx, func(tx *ent.Tx) error {
		return s.service.UseHypeTX(s.ctx, tx, s.user.ID, 30)
	})
	s.Error(err)
	s.Contains(err.Error(), "not enough hype remaining")

	hype, err := s.app.Client().Hype.
		Query().
		Where(hype2.HasUserWith(user.IDEQ(s.user.ID))).
		Only(s.ctx)
	s.NoError(err)
	s.Equal(20, hype.AmountRemaining)
}
