package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kvngho/vimeovideoconverter/api/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewConverterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	r, err := c.ConvertVideo(ctx, &pb.VideoRequest{Type: "admin", Url: ""})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%v", r.Success)
}
