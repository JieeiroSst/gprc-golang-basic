protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service.proto
    
    

protoc --go_out=paths=source_relative:. path/to/file.proto