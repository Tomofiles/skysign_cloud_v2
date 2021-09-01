module flight-operation

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Tomofiles/skysign_cloud_v2/skysign-common v0.0.0-00010101000000-000000000000
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/stretchr/testify v1.7.0
	google.golang.org/genproto v0.0.0-20210212180131-e7f2df4ecc2d
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gorm.io/gorm v1.20.12
)

replace github.com/Tomofiles/skysign_cloud_v2/skysign-common => ../skysign-common
