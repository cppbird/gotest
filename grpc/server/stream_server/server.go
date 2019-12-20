package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net"

// 	"google.golang.org/grpc"

// 	pb "gotest/grpc/proto/stream_proto"
// )

// type StreamService struct{}

// const (
// 	PORT = "9002"
// )

// func main() {
// 	server := grpc.NewServer()
// 	pb.RegisterStreamServiceServer(server, &StreamService{})

// 	lis, err := net.Listen("tcp", ":"+PORT)
// 	if err != nil {
// 		log.Fatalf("net.Listen err: %v", err)
// 	}

// 	server.Serve(lis)
// }

// func (s *StreamService) ServerStream(r *pb.ServerStreamReq, stream pb.StreamService_ServerStreamServer) error {
// 	f, err := ioutil.ReadFile("/Users/jdq/go/src/gotest/fuck.mp4")
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	step := 1 << 20
// 	partition := len(f)/step + 1
// 	log.Printf("step: %d, partition: %d", step, partition)
// 	for n := 0; n < partition; n++ {
// 		var end int
// 		if (n+1)*step < len(f) {
// 			end = (n + 1) * step
// 		} else {
// 			end = len(f)
// 		}
// 		err := stream.Send(
// 			&pb.StreamPoint{
// 				Seq:    int64(n),
// 				BinVal: f[n*step : end],
// 			})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
// func (s *StreamService) ClientStream(stream pb.StreamService_ClientStreamServer) error {
// 	return nil
// }

// func (s *StreamService) BidiStream(stream pb.StreamService_BidiStreamServer) error {
// 	return nil
// }
