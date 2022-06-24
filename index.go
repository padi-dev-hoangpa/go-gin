package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"example/hoangdz/gin/client"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example/hoangdz/gin/square/square"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func setupRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/api/:number", func(c *gin.Context) {
		number := c.Param("number")
		numberInt, err := strconv.ParseInt(number, 0, 32)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  false,
			})
		}
		response := gin.H{}
		for i := 0; i < int(numberInt); i++ {
			result, _ := client.RestGetSquare(int32(i + 1))
			response[strconv.Itoa(i+1)] = result.Result
		}

		c.JSON(http.StatusOK, response)
	})
	r.GET("/rpc/:number", func(c *gin.Context) {
		number := c.Param("number")
		numberInt, err := strconv.ParseInt(number, 0, 32)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  false,
			})
		}
		// Set up a connection to the server.
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		rpcClient := pb.NewSquareRpcClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response := gin.H{}
		for i := 0; i < int(numberInt); i++ {
			result := client.GrpcGetSquare(rpcClient, ctx, int32(i+1))
			response[strconv.Itoa(i+1)] = result
		}

		c.JSON(http.StatusOK, response)
	})
	return r
}

func main() {
	r := setupRoute()
	r.Run(":8080")
}
