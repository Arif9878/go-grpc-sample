generate-grpc-server:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./grpc-server/chat/chat.proto

run-grpc-server:
	go run ./grpc-server/main.go

run-grpc-client:
	go run ./grpc-client/main.go

tidy:
	go mod vendor
