package service

import (
	"context"
	"nevissGo/app/service/mocks"
	"nevissGo/ent"
	"testing"
)

//go:generate mockery --name HypeBridge
type HypeBridge interface {
	UseHypeTX(ctx context.Context, tx *ent.Tx, userID int64, amount int) error
}

type Bridge struct {
	Hype HypeBridge
}

type TestingBridge struct {
	Hype *mocks.HypeBridge

	Bridge
}

func TestBridge(t *testing.T) TestingBridge {
	hype := mocks.NewHypeBridge(t)

	return TestingBridge{
		Hype: hype,
		Bridge: Bridge{
			Hype: hype,
		},
	}
}
