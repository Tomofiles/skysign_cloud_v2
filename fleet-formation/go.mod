module github.com/Tomofiles/skysign_cloud_v2/fleet-formation

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Tomofiles/skysign_cloud_v2/skysign-common v0.0.0-00010101000000-000000000000
	github.com/Tomofiles/skysign_cloud_v2/skysign-proto v0.0.0-00010101000000-000000000000
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.1.2
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.20.12
)

replace github.com/Tomofiles/skysign_cloud_v2/skysign-common => ../skysign-common

replace github.com/Tomofiles/skysign_cloud_v2/skysign-proto => ../skysign-proto
