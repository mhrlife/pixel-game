package framework

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/centrifugal/gocent/v3"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"strings"
)

//go:generate mockery --name Centrifugo
type Centrifugo interface {
	PersonalMessage(ctx context.Context, userID any, eventName string, data any) error
	OnlineWSUsers(ctx context.Context) (int, error)
	PersonalMany(ctx context.Context, usersIds []any, eventName string, data any) error
	Broadcast(ctx context.Context, eventName string, data any) error
}

func EventUserIDs(ids []int64) []any {
	return lo.Map(ids, func(item int64, _ int) any {
		return ids
	})
}

func EventGameIDs(ids []string) []any {
	return lo.Map(ids, func(item string, _ int) any {
		return ids
	})
}

var _ Centrifugo = &CentrifugoClient{}

type CentrifugoClient struct {
	centClient *gocent.Client
}

func NewCentrifugoClient(centClient *gocent.Client) *CentrifugoClient {
	return &CentrifugoClient{centClient: centClient}
}

func (c *CentrifugoClient) Broadcast(ctx context.Context, eventName string, data any) error {
	if c.centClient == nil {
		return nil
	}

	dataBytes, err := json.Marshal(map[string]any{
		"event": eventName,
		"data":  data,
	})
	if err != nil {
		logrus.WithError(err).Error("couldn't marshal broadcast data")
		return err
	}

	_, err = c.centClient.Publish(ctx, "personal:broadcast", dataBytes)
	if err != nil {
		logrus.WithError(err).Error("couldn't publish broadcast message")
		return err
	}
	return err
}

func (c *CentrifugoClient) PersonalMany(ctx context.Context, usersIds []any, eventName string, data any) error {
	if c.centClient == nil {
		return nil
	}

	dataBytes, err := json.Marshal(map[string]any{
		"event": eventName,
		"data":  data,
	})
	if err != nil {
		logrus.WithError(err).Error("couldn't marshal broadcast data")
		return err
	}

	channels := lo.Map[any, string](usersIds, func(item any, _ int) string {
		return fmt.Sprintf("personal:#%d", item)
	})
	_, err = c.centClient.Broadcast(ctx, channels, dataBytes)
	if err != nil {
		logrus.WithError(err).WithField("channels", strings.Join(channels, "-")).Error("couldn't publish personal many message")
		return err
	}
	return err
}

func (c *CentrifugoClient) PersonalMessage(ctx context.Context, userID any, eventName string, data any) error {
	if c.centClient == nil {
		return nil
	}

	dataBytes, err := json.Marshal(map[string]any{
		"event": eventName,
		"data":  data,
	})
	if err != nil {
		logrus.WithError(err).Error("couldn't marshal personal data")
		return err
	}

	_, err = c.centClient.Publish(ctx, fmt.Sprintf("personal:#%v", userID), dataBytes)
	if err != nil {
		logrus.WithError(err).Error("couldn't publish personal message")
		return err
	}
	return err
}

func (c *CentrifugoClient) OnlineWSUsers(ctx context.Context) (int, error) {
	if c.centClient == nil {
		return 0, nil
	}

	info, err := c.centClient.Info(ctx)
	if err != nil {
		logrus.WithError(err).Error("couldn't fetch centrifugo info")
		return 0, err
	}

	return info.Nodes[0].NumClients, nil
}
