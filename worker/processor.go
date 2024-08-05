package worker

import (
	"context"

	"github.com/hibiken/asynq"
	db "github.com/nochzato/go-bank/db/sqlc"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) *RedisTaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (rtp *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, rtp.ProcessTaskSendVerifyEmail)

	return rtp.server.Start(mux)
}
