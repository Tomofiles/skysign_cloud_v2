module github.com/Tomofiles/skysign_cloud_v2/helper-api

go 1.13

require (
	github.com/Tomofiles/skysign_cloud_v2/skysign-common v0.0.0-00010101000000-000000000000
	github.com/Tomofiles/skysign_cloud_v2/skysign-proto v0.0.0-00010101000000-000000000000
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	google.golang.org/grpc v1.40.0
)

replace github.com/Tomofiles/skysign_cloud_v2/skysign-common => ../skysign-common

replace github.com/Tomofiles/skysign_cloud_v2/skysign-proto => ../skysign-proto
