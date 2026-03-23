package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "Distributed-ORDER-system/proto"

	"google.golang.org/grpc"
)

func main() {

	mode := "success" 

	if len(os.Args) > 1 {
		mode = os.Args[1]
	}

	log.Println("Mode:", mode)

	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	var timeout time.Duration

	switch mode {
	case "timeout":
		timeout = 1 * time.Second 
	default:
		timeout = 5 * time.Second 
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()


	if mode == "timeout" {
		log.Println("Simulating delay...")
		time.Sleep(3 * time.Second)
	}

	log.Println("Sending Order Request...")

	res, err := client.CreateOrder(ctx, &pb.OrderRequest{
		UserId: 1,
		Amount: 100,
	})

	if err != nil {
		log.Println("Request FAILED:", err)
		return
	}

	log.Println("Final Result:", res.Status)
}