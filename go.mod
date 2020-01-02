module gotest

go 1.12

require (
	github.com/alecthomas/jsonschema v0.0.0-20191017121752-4bb6e3fae4f2
	github.com/cppbird/cron/v3 v3.0.1
	github.com/cppbird/goSnowFlake v0.0.0-20180906112711-fc763800eec9
	github.com/cppbird/mmh3 v0.0.0-20140820141314-64b85163255b
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.3.2
	github.com/google/wire v0.4.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/json-iterator/go v1.1.8
	github.com/pkg/errors v0.8.1
	github.com/reusee/mmh3 v0.0.0-20140820141314-64b85163255b
	github.com/zheng-ji/goSnowFlake v0.0.0-20180906112711-fc763800eec9
	go-common v0.2.27
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0 // indirect
	golang.org/x/perf v0.0.0-20191209155426-36b577b0eb03 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc v1.23.1
)

replace go-common => git.bilibili.co/platform/go-common v0.2.27
