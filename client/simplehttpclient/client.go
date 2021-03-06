package main

import (
	"context"
	"go-grpc-example/pkg/gtls"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9003"

func main()  {
	tlsClient := gtls.Client{
		ServerName: "go-grpc-example",
		CertFile: "../../conf/server/server.pem",
	}
	c,err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err:%v",err)
	}
	conn,err := grpc.Dial(":"+PORT,grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err:%v",err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp,err := client.Search(context.Background(),&pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Searc err:%v",err)
	}
	log.Printf("resp:%s",resp.GetResponse())
}
