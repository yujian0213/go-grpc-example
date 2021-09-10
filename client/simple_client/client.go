package main

import (
	"context"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9001"

func main()  {
	conn,err := grpc.Dial(":"+PORT,grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc.Dial:%v",err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp,err :=client.Search(context.Background(),&pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalln("client.Search err:%v",err)

	}
	log.Printf("resp:%s",resp.GetResponse())
}
