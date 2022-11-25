package main

import (
	"log"
	"net"

	"github.com/MSalehKoinWorks/grpc-two/controllers/account"
	pb "github.com/MSalehKoinWorks/grpc-two/proto/account"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Server running ...")

	listener, err := net.Listen("tcp", ":6666")

	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDepositServiceServer(grpcServer, &account.DepositServer{})

	log.Fatalln(grpcServer.Serve(listener))
}
