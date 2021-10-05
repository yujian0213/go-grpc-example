package main

import (
	"context"
	"go-grpc-example/pkg/gtls"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9004"

type Auth struct {
	AppKey string
	AppSecret string
}

func (a *Auth) GetRequestMetaData(ctx context.Context,uri ...string) (map[string]string,error)  {
	return map[string]string{"app_key":a.AppKey,"app_secret":a.AppSecret},nil
}
func (a *Auth) RequireTransportSecurity() bool  {
	return true
}
func main()  {
	tlsClient := gtls.Client{
		ServerName: "go-grpc-example",
		CertFile: "../../conf/server/server.pem",
	}
	c,err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err : %v",err)
	}
	auth  := Auth{
		AppSecret: "20210927",
		AppKey: "yujian",
	}
	conn,err := grpc.Dial(":"+PORT,grpc.WithTransportCredentials(c),grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("rpc.Dial err : %v",err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp,err := client.Search(context.Background(),&pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err:%v",err)
	}
	log.Printf("resp:%s",resp.GetResponse())
}
