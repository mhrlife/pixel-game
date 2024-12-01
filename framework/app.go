package framework

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"nevissGo/ent"
	"slices"
)

type Module interface {
	Init(app App)
}

type Config struct {
	Addr string
}

type App struct {
	config Config
	Event  Centrifugo

	client    *ent.Client
	endpoints *Endpoints
	validate  *validator.Validate
}

func NewApp(client *ent.Client, cent Centrifugo, config Config) *App {
	validate := validator.New()

	app := &App{
		config:    config,
		client:    client,
		endpoints: &Endpoints{endpoints: make(map[string]EndpointHandler)},
		validate:  validate,
		Event:     cent,
	}

	if err := validate.RegisterValidation("action", func(fl validator.FieldLevel) bool {
		return slices.Contains(app.endpoints.Actions(), fl.Field().String())
	}); err != nil {
		logrus.WithError(err).Fatal("couldn't register action validation")
	}

	return app

}

func (a *App) TX(ctx context.Context, fn func(tx *ent.Tx) error) error {
	return WithTx(ctx, a.client, fn)
}

func (a *App) Client() *ent.Client {
	return a.client
}

func (a *App) RegisterEndpoints(endpoints ...Endpoint) {
	for _, endpoint := range endpoints {
		endpoint.Endpoints(a.endpoints)
	}
}

type CallRequest struct {
	Action string `validate:"required,action"`
}

func (a *App) ServeEndpoints() error {
	e := echo.New()

	e.Use(middleware.Recover())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := ExtErrorCode(err)
		message := ExtErrorMessage(err)
		fields := ExtErrorFields(err)

		response := map[string]any{
			"ok":         false,
			"error_code": code,
			"message":    message,
		}

		if fields != nil {
			response["fields"] = fields
		}

		c.JSON(200, response)
	}

	for _, middlewareFunc := range a.endpoints.middlewares {
		e.Use(middlewareFunc)
	}

	e.POST("/call", func(c echo.Context) error {
		user := c.Get("user").(ent.User)

		ctx := &Context{
			Context: c,
			App:     a,
			User:    &user,
		}

		request, err := BindAndValidate[CallRequest](ctx)
		if err != nil {
			return err
		}

		return a.endpoints.endpoints[request.Action](ctx)
	})

	return e.Start(a.config.Addr)
}
