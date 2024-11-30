package framework

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
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

	client    *ent.Client
	endpoints *Endpoints
	validate  *validator.Validate
}

func NewApp(client *ent.Client, config Config) *App {
	validate := validator.New()

	app := &App{
		config:    config,
		client:    client,
		endpoints: &Endpoints{endpoints: make(map[string]EndpointHandler)},
		validate:  validate,
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

		// if error as validator.ValidationErrors
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			c.JSON(400, map[string]any{
				"fields": lo.Map(errs, func(err validator.FieldError, _ int) string {
					return err.Error()
				}),
			})
			return
		}

		logrus.WithError(err).Error("internal server error")

		c.String(500, "internal server error")
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
