package endpoint

import (
	"nevissGo/app/service"
	"nevissGo/framework"
)

var _ framework.Endpoint = &OnlineUsers{}

type OnlineUsers struct {
	service *service.OnlineUsers
}

func NewOnlineUsers(service *service.OnlineUsers) *OnlineUsers {
	return &OnlineUsers{
		service: service,
	}
}

func (e *OnlineUsers) Endpoints(router *framework.Endpoints) {
	router.Register("online_users/count", e.GetOnlineUsersCount)
}

func (e *OnlineUsers) GetOnlineUsersCount(c *framework.Context) error {
	count, err := e.service.GetOnlineUsersCount(c.Request().Context())
	if err != nil {
		return err
	}
	
	return c.Ok(count)
}
