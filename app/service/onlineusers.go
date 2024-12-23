package service

import (
	"context"
	"nevissGo/framework"
)

type OnlineUsers struct {
	app *framework.App
}

func NewOnlineUsers(app *framework.App) *OnlineUsers {
	return &OnlineUsers{
		app: app,
	}
}

func (s *OnlineUsers) GetOnlineUsersCount(ctx context.Context) (int, error) {
	result, err := s.app.Event.OnlineWSUsers(ctx)
	if err != nil {
		return 0, framework.NewInternalError("Failed to get online users count")
	}
	return result, nil
}
