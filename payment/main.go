package main

import (
	"context"
	"log"
	"net"

	pb "Distributed-ORDER-system/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) CreatePayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Println("Payment processed")
	return &pb.PaymentResponse{Status: "SUCCESS"}, nil
}

func (s *server) RefundPayment(ctx context.Context, req *pb.RefundRequest) (*pb.RefundResponse, error) {
	log.Println("Payment refunded")
	return &pb.RefundResponse{Status: "REFUNDED"}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()

	pb.RegisterPaymentServiceServer(s, &server{})

	log.Println("Payment Service running on 50051")
	s.Serve(lis)
}