package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"nevissGo/ent"
	"nevissGo/framework"
)

type UsersSuite struct {
	suite.Suite
	app     *framework.TestingApp
	service *Users
	ctx     context.Context
}

func TestUsers(t *testing.T) {
	suite.Run(t, new(UsersSuite))
}

func (s *UsersSuite) SetupTest() {
	s.app = framework.NewTestingApp(s.T())
	s.service = NewUsers(s.app.App)
	s.ctx = context.Background()
}

func (s *UsersSuite) TestGetOrRegister() {
	user := &ent.User{
		ID:          1,
		DisplayName: "1",
	}
	err := s.service.GetOrRegister(s.ctx, user)
	s.NoError(err)

	createdUser, err := s.app.Client().User.Get(s.ctx, user.ID)
	s.NoError(err)
	s.Equal(user.ID, createdUser.ID)
	s.Equal(user.DisplayName, createdUser.DisplayName)
	s.Require().NotEmpty(createdUser.GameID)

	user.DisplayName = "3"
	err = s.service.GetOrRegister(s.ctx, user)
	s.NoError(err)

	updatedUser, err := s.app.Client().User.Get(s.ctx, user.ID)
	s.NoError(err)
	s.Equal(user.ID, updatedUser.ID)
	s.Equal("1", updatedUser.DisplayName)
}
