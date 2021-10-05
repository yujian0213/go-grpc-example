package main

import (
	"context"
	"go-grpc-example/pkg/gtls"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type Auth struct {
	AppKey string
	AppSecret string
}
type SearchService struct {
	auth *Auth
}

func (s *SearchService) Search(ctx context.Context,r *pb.SearchRequest) (*pb.SearchResponse,error)  {
	if err := s.auth.Check(ctx);err != nil {
		return nil,err
	}
	return &pb.SearchResponse{Response: r.GetRequest() + "Token Server"},nil
}

const PORT = "9004"

func main()  {
	certFile := "../../conf/server/server.pem"
	keyFile := "../../conf/server/server.key"
	tlsServer := gtls.Server{
		CertFile: certFile,
		KeyFile: keyFile,
	}
	c,err := tlsServer.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err :%v",err)
	}
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server,&SearchService{})
	lis,err := net.Listen("tcp",":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v",err)
	}
	server.ServeHTTP(lis)
}
func (a * Auth) Check(ctx context.Context) error  {
	md,ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated,"metadata.FromIncomingContext err ")
	}
	var (
		appKey string
		appSecret  string
	)
	if value,ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value,ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appKey != a.GetAppKey() || appSecret != a.AppSecret {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}
	return nil
}
func (a *Auth) GetAppKey() string {
	return "yujian"
}

func (a *Auth) GetAppSecret() string {
	return "20210927"
}