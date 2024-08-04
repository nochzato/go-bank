package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/nochzato/go-bank/api"
	db "github.com/nochzato/go-bank/db/sqlc"
	"github.com/nochzato/go-bank/gapi"
	"github.com/nochzato/go-bank/pb"
	"github.com/nochzato/go-bank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to the db: %v", err)
	}

	store := db.NewStore(conn)

	runGrpcServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGoBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %v", err)
	}
}
