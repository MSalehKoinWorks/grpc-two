package account

import (
	"context"
	"fmt"

	pb "github.com/MSalehKoinWorks/grpc-two/proto/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var amount float32 = 0

type DepositServer struct {
	pb.UnimplementedDepositServiceServer
}

func (*DepositServer) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	if req.GetAmount() < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot deposit %v", req.GetAmount())
	}

	amount += req.GetAmount()
	fmt.Printf("new deposit: %v, balance: %v\n", req.GetAmount(), amount)

	return &pb.DepositResponse{Ok: true}, nil
}
