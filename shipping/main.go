package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "Distributed-ORDER-system/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedShippingServiceServer
}

func (s *server) StartShipping(ctx context.Context, req *pb.ShippingRequest) (*pb.ShippingResponse, error) {

	if rand.Intn(2) == 0 {
		log.Println("Shipping FAILED")
		return &pb.ShippingResponse{Status: "FAILED"}, nil
	}

	log.Println("Shipping SUCCESS")
	return &pb.ShippingResponse{Status: "SUCCESS"}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50052")
	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &server{})

	log.Println("Shipping Service running on 50052")
	s.Serve(lis)
}