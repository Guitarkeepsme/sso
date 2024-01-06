package app

import "google.golang.org/grpc"

type App struct {
	gRPCServer *grpc.Server
}
