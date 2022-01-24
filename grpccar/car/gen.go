package car

//go:generate -command compile_proto protoc -I../proto
//go:generate compile_proto car.proto --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:.
