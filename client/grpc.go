package client

import (
	"context"
	"flag"
	"log"

	pb "example/hoangdz/gin/square/square"
)

func GrpcGetSquare(c pb.SquareRpcClient, ctx context.Context, number int32) int32 {
	flag.Parse()
	r, err := c.FindSquareUnary(ctx, &pb.Input{Number: number})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetResult()
}
