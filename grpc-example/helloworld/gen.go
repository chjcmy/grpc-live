package helloworld

//go:generate -command compile_proto protoc -I../protos
//go:generate compile_proto helloworld.proto --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:.
