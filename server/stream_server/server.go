package main

import (
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type StreamService struct {

}

const PORT = "9002"

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for i := 0; i <7 ; i++ {
		err := stream.Send(&pb.StreamResponse{Pt: &pb.StreamPoint{Value:r.Pt.Value + int32(i) ,Name:r.Pt.Name }})
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	return nil
}
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	return nil
}


func main()  {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server,&StreamService{})
	lis,err := net.Listen("tcp",":"+PORT)
	if err != nil {
		log.Fatalln("net.Listen:%v",err)
	}
	server.Serve(lis)
}