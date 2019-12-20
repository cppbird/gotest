package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// )

// func main() {
// 	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure(), grpc.WithChainUnaryInterceptor(grpc_retry.UnaryClientInterceptor()))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer conn.Close()
// 	c := proto.NewTestClient(conn)
// 	_, err = c.TestRpc(context.TODO(), &proto.TestRpcReq{Tp: 123}, grpc_retry.WithMax(4), grpc_retry.WithPerRetryTimeout(1000*time.Millisecond), grpc_retry.WithCodes(codes.InvalidArgument))
// 	fmt.Println(err)
// }
