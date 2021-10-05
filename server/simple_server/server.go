package main

import (
	"context"
	"go-grpc-example/pkg/gtls"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"runtime/debug"
)

type SearchService struct {

}


func ( s *SearchService ) Search(ctx context.Context,r *pb.SearchRequest) (*pb.SearchResponse,error) {
	return &pb.SearchResponse{Response: r.GetRequest() + "server"},nil
}

const PORT = "9001"

func main()  {
	tlsServer := gtls.Server{
		CaFile: "../../conf/ca.pem",
		CertFile: "../../conf/server/server.pem",
		KeyFile: "../../conf/server/server.key",
	}
	c,err := tlsServer.GetCredentialsByCA()
	if err != nil {
		log.Printf("GetCredentialsByCA err:%v",err)
	}
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,
			LoggingInterceptor,
			),
	}
	server := grpc.NewServer(opts...)

	pb.RegisterSearchServiceServer(server,&SearchService{})
	lis,err := net.Listen("tcp",":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen:%v\n", err)
	}
	server.Serve(lis)
}
func LoggingInterceptor(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler) (interface{},error) {
	log.Printf("grpc method: %s,%v",info.FullMethod,req)
	resp,err := handler(ctx,req)
	log.Printf("grpc method: %s,%v",info.FullMethod,req)
	return resp,err
}
func RecoveryInterceptor(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler) (resp interface{},err error) {
	defer func() {
		if e:= recover();e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal,"Panic err:%v",e)
		}
	}()
	return handler(ctx,req)
}