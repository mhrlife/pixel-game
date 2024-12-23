package service

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mhrlife/tonference/internal/ent"
	"github.com/mhrlife/tonference/pkg/framework"
	"github.com/stretchr/testify/require"
	"github.com/teris-io/shortid"
	"testing"
)

func NewTestingService(t *testing.T) *Service {
	fileID := shortid.MustGenerate()
	client, err := ent.Open("sqlite3", "file:"+fileID+"?mode=memory&cache=shared&_fk=1")

	require.NoError(t, err)
	require.NoError(t, client.Schema.Create(context.Background()))

	app := framework.NewApp(client, nil, framework.Config{})

	return NewService(client, app, nil)
}
