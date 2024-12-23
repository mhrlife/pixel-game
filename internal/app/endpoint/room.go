package endpoint

import (
	"github.com/mhrlife/tonference/internal/app/serializer"
	"github.com/mhrlife/tonference/internal/app/service"
	"github.com/mhrlife/tonference/pkg/framework"
)

var _ framework.Endpoint = &Rooms{}

type Rooms struct {
	service *service.Service
}

func NewRooms(service *service.Service) *Rooms {
	return &Rooms{
		service: service,
	}
}

func (e *Rooms) Endpoints(router *framework.Endpoints) {
	router.Register("rooms/create", e.Create)
	router.Register("rooms/mine", e.MyRooms)
	router.Register("rooms/delete", e.DeleteRoom)
	router.Register("rooms/token", e.Token)
}

func (e *Rooms) Create(c *framework.Context) error {
	createdRoom, err := e.service.CreateRoom(c.Request().Context(), service.CreateRoomDto{
		CreatorID: c.User.ID,
	})
	if err != nil {
		return err
	}

	return c.Ok(serializer.NewRoomSerializer(createdRoom))
}

func (e *Rooms) MyRooms(c *framework.Context) error {
	myRooms, err := e.service.MyRooms(c.Request().Context(), c.User.ID)
	if err != nil {
		return err
	}

	return c.Ok(serializer.List(myRooms, serializer.NewRoomSerializer))
}

type DeleteRoomRequest struct {
	ID string `json:"id" validate:"required"`
}

func (e *Rooms) DeleteRoom(c *framework.Context) error {
	request, err := framework.BindAndValidate[DeleteRoomRequest](c)
	if err != nil {
		return err
	}

	err = e.service.DeleteRoom(c.Request().Context(), request.ID, c.User.ID)
	if err != nil {
		return err
	}

	return c.Ok("deleted successfully")
}

type TokenRequest struct {
	ID string `json:"id" validate:"required"`
}

func (e *Rooms) Token(c *framework.Context) error {
	request, err := framework.BindAndValidate[TokenRequest](c)
	if err != nil {
		return err
	}

	token, err := e.service.GenerateRoomToken(c.Request().Context(), request.ID, c.User)
	if err != nil {
		return err
	}

	return c.Ok(token)
}
