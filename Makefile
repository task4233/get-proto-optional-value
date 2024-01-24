proto/gen:
	protoc  --go_out=proto --go-grpc_out=require_unimplemented_servers=false:proto proto/color.proto

test:
	go test -race -shuffle=on ./...
