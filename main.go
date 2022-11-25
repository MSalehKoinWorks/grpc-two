package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/MSalehKoinWorks/grpc-two/proto"
	"github.com/MSalehKoinWorks/grpc-two/server"
)

func main() {
	log.Println("Server running ...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDepositServiceServer(grpcServer, &server.DepositServer{})

	log.Fatalln(grpcServer.Serve(listener))
}
