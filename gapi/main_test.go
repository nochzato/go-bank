package gapi

import (
	"testing"
	"time"

	db "github.com/nochzato/go-bank/db/sqlc"
	"github.com/nochzato/go-bank/util"
	"github.com/nochzato/go-bank/worker"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	t.Helper()

	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
