package main

import (
	"fmt"
	"github.com/artbegolli/spec-gen/api/spec-gen/v1alpha1"
	"google.golang.org/grpc"
	"log"
	"net"
)

// main start a gRPC server and waits for connection
func main() {
	fmt.Println("starting up service")
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := v1alpha1.Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	v1alpha1.RegisterPingServer(grpcServer, &s)
	v1alpha1.RegisterSpecGenServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
