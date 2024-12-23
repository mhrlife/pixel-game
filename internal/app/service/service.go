package service

import (
	"context"
	"fmt"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/mhrlife/tonference/internal/ent"
	"github.com/mhrlife/tonference/pkg/framework"
)

type Service struct {
	client             *ent.Client
	app                *framework.App
	liveKitRoomService *lksdk.RoomServiceClient
}

func NewService(
	client *ent.Client,
	app *framework.App,
	liveKitRoomService *lksdk.RoomServiceClient,
) *Service {
	return &Service{
		client:             client,
		app:                app,
		liveKitRoomService: liveKitRoomService,
	}
}

func (s *Service) Close() {
	s.client.Close()
}

func (s *Service) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
