package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

const PORT = "9001"

func main()  {
	//c,err := credentials.NewClientTLSFromFile("../../conf/server.pem","go-grpc-example")
	//if err != nil {
	//	log.Fatalf("credentials.NewClientTLSFromFile err: %v",err)
	//}
	cert,err := tls.LoadX509KeyPair("../../conf/client/client.pem","../../conf/client//client.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err:%s",err)
	}
	certPool := x509.NewCertPool()
	ca,err := ioutil.ReadFile("../../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err:%s",err)
	}
	if ok:= certPool.AppendCertsFromPEM(ca); !ok  {
		log.Fatalf("certPloo.AppendCertsFromPEM err:%s",err)
	}
	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName: "go-grpc-example",
		RootCAs: certPool,
	})
	conn,err := grpc.Dial(":"+PORT,grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp,err :=client.Search(context.Background(),&pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err:%v\n", err)

	}
	log.Printf("resp:%s",resp.GetResponse())
}


