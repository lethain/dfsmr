
server:
	go run dfsmr_server/main.go

client:
	go run dfsmr_client/main.go

proto:
	protoc -I dfsmr/ dfsmr/dfsmr.proto --go_out=plugins=grpc:dfsmr
