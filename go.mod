module gotest

go 1.12

require (
	github.com/cppbird/cron/v3 v3.0.1
	github.com/gogo/protobuf v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/json-iterator/go v1.1.8
	github.com/pkg/errors v0.8.1
	go-common v0.2.27
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0 // indirect
	google.golang.org/grpc v1.23.1
)

replace go-common => git.bilibili.co/platform/go-common v0.2.27
