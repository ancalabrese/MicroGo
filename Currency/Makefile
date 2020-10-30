.PHONY: protos

protos: 
	protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency

run: 
	echo "Starting up Currency service"
	go run ./main.go

	