module edge-px4

go 1.13

require (
	github.com/Tomofiles/skysign_cloud_v2/skysign-proto v0.0.0-00010101000000-000000000000
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/protobuf v1.5.2
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/Tomofiles/skysign_cloud_v2/skysign-proto => ../skysign-proto
