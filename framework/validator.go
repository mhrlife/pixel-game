package framework

import (
	"github.com/sirupsen/logrus"
)

func BindAndValidate[T any](ctx *Context) (T, error) {
	var t T

	if err := ctx.readInput(); err != nil {
		logrus.WithError(err).Error("couldn't read input")
		return t, err
	}

	if err := ctx.Bind(&t); err != nil {
		return t, err
	}

	if err := ctx.App.validate.Struct(t); err != nil {
		return t, err
	}

	return t, nil
}
