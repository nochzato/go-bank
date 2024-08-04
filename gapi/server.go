package gapi

import (
	"fmt"

	db "github.com/nochzato/go-bank/db/sqlc"
	"github.com/nochzato/go-bank/pb"
	"github.com/nochzato/go-bank/token"
	"github.com/nochzato/go-bank/util"
)

// Server serves HTTP requests.
type Server struct {
	pb.UnimplementedGoBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
