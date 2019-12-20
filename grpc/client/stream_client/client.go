package main

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	"google.golang.org/grpc"

// 	pb "gotest/grpc/proto/stream_proto"
// )

// const (
// 	PORT = "9002"
// )

// func main() {
// 	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("grpc.Dial err: %v", err)
// 	}

// 	defer conn.Close()

// 	client := pb.NewStreamServiceClient(conn)

// 	var stream pb.StreamService_ServerStreamClient

// 	if stream, err = client.ServerStream(context.TODO(), &pb.ServerStreamReq{
// 		Fn: "file",
// 	}); err != nil {
// 		log.Fatalf("client.ServerStream err: %v", err)
// 	}
// 	fileByte := make(map[int64][]byte)
// 	for {
// 		v := &pb.StreamPoint{}
// 		if v, err = stream.Recv(); err != nil {
// 			fmt.Printf("stream.Recv() err :%v", err)
// 			break
// 		}
// 		fileByte[v.Seq] = v.BinVal
// 		log.Printf("recv data, seq: %d", v.Seq)
// 	}
// 	var fileArr []byte
// 	for i := 0; i < len(fileByte); i++ {
// 		fileArr = append(fileArr, fileByte[int64(i)]...)
// 	}
// 	err = ioutil.WriteFile("./fuck.mp4", fileArr, 0644)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
