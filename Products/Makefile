check_swagger:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
	
swagger: check_swagger
	GO111MODULE=off swagger generate spec -o ./api-doc/swagger.yaml --scan-models

run: 
	go run main.go
