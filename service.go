package main

import (
	"context"
	"fmt"
	pb "github.com/shinshin8/golang-grpc-protobuf/gen/go/protobuf"
	"google.golang.org/protobuf/types/known/emptypb"
)


type service struct {
	pb.UnimplementedGrpcServiceServer
}

func NewService() pb.GrpcServiceServer {
	return &service{
		UnimplementedGrpcServiceServer: pb.UnimplementedGrpcServiceServer{},
	}
}

func (s *service) FindEmployee(ctx context.Context, req *pb.FindEmployeeRequest) (*pb.FindEmployeeResponse, error) {
	id := req.GetID()

	if id == 1 {
		return &pb.FindEmployeeResponse{
			Enployee: &pb.Employee{
				ID: 1,
				Name: "john",
				Status: pb.StatusType_STAFF,
			},
		}, nil
	}
	return nil, fmt.Errorf("cannnot find employee")
}

func (s *service) ListEmployee(context.Context, *emptypb.Empty) (*pb.ListEmployeeResponse, error) {
	return &pb.ListEmployeeResponse{
		Employees: []*pb.Employee{
			{
				ID: 1,
				Name: "john",
				Status: pb.StatusType_STAFF,
			},
			{
				ID: 2,
				Name: "emily",
				Status: pb.StatusType_MANAGER,
			},
			{
				ID: 3,
				Name: "mike",
				Status: pb.StatusType_PART_TIME,
			},
		},
	}, nil
}