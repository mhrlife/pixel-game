package framework

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"github.com/teris-io/shortid"
	"nevissGo/ent"
	"testing"
)

type TestingApp struct {
	*App
}

func NewTestingApp(t *testing.T) *TestingApp {
	fileID := shortid.MustGenerate()
	client, err := ent.Open("sqlite3", "file:"+fileID+"?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)

	app := &App{
		client: client,
	}

	require.NoError(t, client.Schema.Create(context.Background()))

	return &TestingApp{App: app}
}
