package main

import (
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9002"

func main()  {
	conn,err := grpc.Dial(":"+PORT,grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc.Dial:%v",err)
	}
	defer conn.Close()
	client := pb.NewStreamServiceClient(conn)
	err = printLists(client,&pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRpc stream Client:List",Value: 2021}})
	if err != nil {
		log.Fatalln("printLists.err: %v",err)
	}
	err = printRecord(client,&pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRpc stream Client:Record",Value: 2021}})
	if err != nil {
		log.Fatalln("printRecord.err: %v",err)
	}
	err = printRoute(client,&pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRpc stream Client:Route",Value: 2021}})
	if err != nil {
		log.Fatalln("printRoute.err: %v",err)
	}

}
func printLists(client pb.StreamServiceClient,r *pb.StreamRequest) error{
	return nil
}

func printRecord(client pb.StreamServiceClient,r *pb.StreamRequest) error{
	return nil
}
func printRoute(client pb.StreamServiceClient,r *pb.StreamRequest) error{
	return nil
}
