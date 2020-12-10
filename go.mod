module github.com/pubgo/golugtest

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.1.5+incompatible
	github.com/dgraph-io/badger/v2 v2.2007.2
	github.com/gofiber/fiber/v2 v2.2.3
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/json-iterator/go v1.1.10
	github.com/pubgo/dix v0.1.3
	github.com/pubgo/golug v0.0.2
	github.com/pubgo/golugin v0.0.0-20201210031001-004b80370754
	github.com/pubgo/xerror v0.3.1
	github.com/pubgo/xlog v0.0.10
	github.com/pubgo/xprocess v0.0.8
	github.com/spf13/cobra v1.1.1
	github.com/twmb/murmur3 v1.1.5
	go.uber.org/atomic v1.7.0
	google.golang.org/genproto v0.0.0-20201204160425-06b3db808446
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
	xorm.io/xorm v1.0.5
)
