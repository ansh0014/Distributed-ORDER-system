package main

import (
	"context"
	"log"
	"time"

	pb "Distributed-ORDER-system/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50053", grpc.WithInsecure())
	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, _ := client.CreateOrder(ctx, &pb.OrderRequest{
		UserId: 1,
		Amount: 100,
	})

	log.Println("Final Result:", res.Status)
}