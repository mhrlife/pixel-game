package service

import (
	"context"
	"testing"

	"github.com/mhrlife/tonference/internal/ent"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite

	*Service
	ctx context.Context
}

func TestUserSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetupSuite() {
	s.Service = NewTestingService(s.T())
	s.ctx = context.Background()
}

func (s *UserTestSuite) TearDownSuite() {
	s.Service.Close()
}

func (s *UserTestSuite) TestGetOrRegister() {
	user := &ent.User{
		ID:          1,
		DisplayName: "1",
	}
	err := s.GetOrRegister(s.ctx, user)
	s.NoError(err)

	createdUser, err := s.app.Client().User.Get(s.ctx, user.ID)
	s.NoError(err)
	s.Equal(user.ID, createdUser.ID)
	s.Equal(user.DisplayName, createdUser.DisplayName)
	s.Require().NotEmpty(createdUser.GameID)

	user.DisplayName = "3"
	err = s.GetOrRegister(s.ctx, user)
	s.NoError(err)

	updatedUser, err := s.app.Client().User.Get(s.ctx, user.ID)
	s.NoError(err)
	s.Equal(user.ID, updatedUser.ID)
	s.Equal("1", updatedUser.DisplayName)
}
