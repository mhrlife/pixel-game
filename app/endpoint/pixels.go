// app/endpoint/pixels.go
package endpoint

import (
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
	"nevissGo/app/serializer"
	"nevissGo/app/service"
	"nevissGo/ent/user"
	"nevissGo/framework"
)

var _ framework.Endpoint = &Pixels{}

type Pixels struct {
	service *service.Pixels
}

func NewPixels(service *service.Pixels) *Pixels {
	return &Pixels{
		service: service,
	}
}

func (p *Pixels) Endpoints(router *framework.Endpoints) {
	router.Register("pixels/update", p.UpdatePixel)
	router.Register("pixels/board", p.GetBoard)
}

type UpdatePixelDto struct {
	PixelID  int    `json:"pixel_id" validate:"required"`
	NewColor string `json:"new_color" validate:"required"`
}

func (p *Pixels) UpdatePixel(c *framework.Context) error {
	request, err := framework.BindAndValidate[UpdatePixelDto](c)
	if err != nil {
		return eris.Wrap(err, "failed to bind and validate request")
	}
	err = p.service.UpdateColor(c.Request().Context(), request.PixelID, request.NewColor, c.User.ID)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"pixel_id":  request.PixelID,
			"user_id":   user.ID,
			"new_color": request.NewColor,
		}).Error("Failed to update pixel color")
		return eris.Wrap(err, "failed to update pixel color")
	}
	return c.Ok("Pixel color updated successfully")
}

func (p *Pixels) GetBoard(c *framework.Context) error {
	board, err := p.service.GetBoard(c.Request().Context())
	if err != nil {
		return eris.Wrap(err, "failed to get board")
	}
	return c.Ok(serializer.NewBoard(board))
}
