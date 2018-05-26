
server:
	go build -o dsrv dfsmr_server/main.go

test_server:
	go test ./dfsmr_server

client:
	go build -o dcli dfsmr_client/main.go

test_client:
	go test ./dfsmr_client

proto:
	protoc -I dfsmr/ dfsmr/dfsmr.proto --go_out=plugins=grpc:dfsmr
