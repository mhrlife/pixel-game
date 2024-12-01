package endpoint

import (
	"github.com/rotisserie/eris"
	"nevissGo/app/serializer"
	"nevissGo/app/service"
	"nevissGo/framework"
)

var _ framework.Endpoint = &Hype{}

type Hype struct {
	service *service.Hype
}

func NewHype(service *service.Hype) *Hype {
	return &Hype{
		service: service,
	}
}

func (h *Hype) Endpoints(router *framework.Endpoints) {
	router.Register("hype/count", h.GetHype)
}

func (h *Hype) GetHype(c *framework.Context) error {
	hype, err := h.service.GetHype(c.Request().Context(), c.User.ID)
	if err != nil {
		return eris.Wrap(err, "failed to get hype")
	}
	return c.Ok(serializer.NewHype(hype))
}
