package main

import (
	"context"
	"go-bbq/pkg/net/grpc/go-grpc-middleware/retry"
	"gotest/proto"
	"time"

	"google.golang.org/grpc/codes"
)

func main() {

	_, err = cli.TestRpc(
		context.Background(),
		&proto.TestRpcReq{},
		retry.WithMax(3),
		retry.WithTimeoutRetry(true),
		retry.WithPerRetryTimeout(5*time.Millisecond),
		retry.WithCodes(codes.Unknown),
	)
	// s := grpc.NewServer()
	// svc := &Service{}
	// pb.RegisterTestServer(s, svc)
	// lis, err := net.Listen("tcp",
	// 	"127.0.0.1:9000")
	// fmt.Println(err)
	// if err = s.Serve(lis); err != nil {
	// 	fmt.Println("fuck lis failed")
	// 	return
	// }

	// type Service struct {
	// }

	// func (s *Service) TestRpc(context.Context, *pb.TestRpcReq) (*pb.TestRpcResp, error) {
	// 	fmt.Println(time.Now())
	// 	return &pb.TestRpcResp{}, status.Error(codes.InvalidArgument, "call failed")
}
