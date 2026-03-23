package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "Distributed-ORDER-system/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *server) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {

	log.Println("Order Created")

	paymentConn, _ := grpc.NewClient("payment:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	paymentClient := pb.NewPaymentServiceClient(paymentConn)


	shippingConn, _ := grpc.NewClient("shipping:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	shippingClient := pb.NewShippingServiceClient(shippingConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	paymentRes, _ := paymentClient.CreatePayment(ctx, &pb.PaymentRequest{
		OrderId: 1,
		Amount:  req.Amount,
	})

	if paymentRes.Status != "SUCCESS" {
		log.Println("Payment Failed")
		return &pb.OrderResponse{Status: "FAILED"}, nil
	}

	shippingRes, _ := shippingClient.StartShipping(ctx, &pb.ShippingRequest{
		OrderId: 1,
	})

	if shippingRes.Status != "SUCCESS" {
		log.Println("Shipping Failed → Refund")

		paymentClient.RefundPayment(ctx, &pb.RefundRequest{
			OrderId: 1,
		})

		return &pb.OrderResponse{Status: "FAILED"}, nil
	}

	log.Println("Order Completed")
	return &pb.OrderResponse{Status: "SUCCESS"}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50053")
	s := grpc.NewServer()

	pb.RegisterOrderServiceServer(s, &server{})

	log.Println("Order Service running on 50053")
	s.Serve(lis)
}