package framework

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"nevissGo/ent"
)

type Context struct {
	echo.Context
	App  *App
	User *ent.User

	input map[string]any
}

func (c *Context) readInput() error {
	if c.input != nil {
		return nil
	}

	c.input = make(map[string]any)

	if err := c.Context.Bind(&c.input); err != nil {
		return err
	}

	return nil
}

func (c *Context) Bind(val any) error {
	bytes, err := json.Marshal(c.input)
	if err != nil {
		logrus.WithError(err).Error("couldn't marshal input")
		return err
	}

	if err := json.Unmarshal(bytes, val); err != nil {
		logrus.WithError(err).Error("couldn't unmarshal input")
		return err
	}

	return nil
}

func (c *Context) Error(errorCode int, message any) error {
	return c.JSON(200, map[string]any{
		"ok":         false,
		"message":    message,
		"error_code": errorCode,
	})
}

func (c *Context) Ok(data any) error {
	return c.JSON(200, map[string]any{
		"ok":   true,
		"data": data,
	})
}
