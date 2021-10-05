package main

import (
	"context"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

const PORT = "9002"

func main()  {
	conn,err := grpc.Dial(":"+PORT,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewStreamServiceClient(conn)
	err = printLists(client,&pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRpc stream Client:List",Value: 2021}})
	if err != nil {
		log.Fatalf("printLists.err: %v\n", err)
	}
	err = printRecord(client,&pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRpc stream Client:Record",Value: 2021}})
	if err != nil {
		log.Fatalf("printRecord.err: %v\n", err)
	}
	err = printRoute(client,&pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRpc stream Client:Route",Value: 2021}})
	if err != nil {
		log.Fatalf("printRoute.err: %v\n", err)
	}

}
func printLists(client pb.StreamServiceClient,r *pb.StreamRequest) error{
	stream,err := client.List(context.Background(),r)
	if err != nil {
		return err
	}
	for  {
		resp,err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp:pt.name:%s,pt.value:%d",resp.Pt.Name,resp.Pt.Value)
	}
	return nil
}

func printRecord(client pb.StreamServiceClient,r *pb.StreamRequest) error{
	stream,err := client.Record(context.Background())
	if err != nil{
		return err
	}
	for i := 0; i < 6; i++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}
	resp,err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("resp:pt.name:%s,pt.value:%d",resp.Pt.Name,resp.Pt.Value)
	return nil
}
func printRoute(client pb.StreamServiceClient,r *pb.StreamRequest) error{
	stream,err := client.Route(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n <=6 ; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}
		resp,err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("resp:pt.name:%s,pt.value:%d",resp.Pt.Name,resp.Pt.Value)

	}
	_ = stream.CloseSend()
	return nil
}
