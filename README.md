# Go GRPC Protobuf Example


## Generate proto and grpc codes
 protoc --go_out=pb --go_opt=paths=source_relative \
     --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
     ./person.proto