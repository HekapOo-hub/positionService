package main

import (
	"context"
	"github.com/HekapOo-hub/positionService/internal/handler"
	"github.com/HekapOo-hub/positionService/internal/proto/positionpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	positionHandler, err := handler.NewPositionHandler(context.Background())
	if err != nil {
		return
	}
	lis, err := net.Listen("tcp", handler.PositionPort)
	if err != nil {
		log.Warnf("error %v", err)
		return
	}
	server := grpc.NewServer()
	positionpb.RegisterPositionServiceServer(server, positionHandler)
	if err = server.Serve(lis); err != nil {
		log.Warnf("%v", err)
	}
}
