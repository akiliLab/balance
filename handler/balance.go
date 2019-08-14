package handler

import (
	"context"

	pb "github.com/akiliLab/balancee/proto"
)

// BalanceServiceServer : server for Balance
type BalanceServiceServer struct {
}

var (
	balance *pb.Balance
)

// Balance : Get Balance
func (s *BalanceServiceServer) Balance(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceResponse, error) {
	balance := pb.Balance{
		Balance:      5000,
		TotalBalance: 6000,
		Currency:     "USD",
		SpendToday:   50,
	}

	// CallGrpcService(ctx, "transcation:50051")
	// CallGrpcService(ctx, "balance:50051")

	return &pb.BalanceResponse{
		Balance: &balance,
	}, nil
}
