package main

import (
	"go-grpc/proto/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()

	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter a"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter b"})
			return
		}

		request := &proto.Request{A: int64(a), B: int64(b)}

		response, err := client.Add(ctx, request)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while add function"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": response})

	})

	g.GET("/multiply/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter a"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter b"})
			return
		}

		request := &proto.Request{A: int64(a), B: int64(b)}

		response, err := client.Multiply(ctx, request)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while multiply function"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": response})
	})

	if err := g.Run(":8081"); err != nil {
		panic(err)
	}
}
