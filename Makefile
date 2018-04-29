
server:
	go build -o dsrv dfsmr_server/main.go 

client:
	go build -o dcli dfsmr_client/main.go

proto:
	protoc -I dfsmr/ dfsmr/dfsmr.proto --go_out=plugins=grpc:dfsmr
