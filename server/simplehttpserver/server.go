package main

import (
	"context"
	"go-grpc-example/pkg/gtls"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strings"
)

type SearchService struct {}

func (s *SearchService) Search(ctx context.Context,r *pb.SearchRequest) (*pb.SearchResponse,error) {
	return &pb.SearchResponse{Response: r.GetRequest() + "HTTP Server"},nil
}

const PORT = "9003"

func main()  {
	certFile := "../../conf/server/server.pem"
	keyFile := "../../conf/server/server.key"
	tlsServer := gtls.Server{
		CertFile: certFile,
		KeyFile: keyFile,
	}
	c,err := tlsServer.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err:%v",err)
	}
	mux := GetHTTPServerNux()
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server,&SearchService{})
	http.ListenAndServeTLS(":"+ PORT,
		certFile,
		keyFile,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"),"application/grpc"){
				server.ServeHTTP(w,r)
			}else{
				mux.ServeHTTP(w,r)
			}
		}),
		)
}
func GetHTTPServerNux() *http.ServeMux  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yujian:go-grpc-example"))
	})
	return mux
}
