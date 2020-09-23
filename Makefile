# check swagger
check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

# create swagger docs
swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./swagger.json --scan-models

# serve swagger
serve-swagger: check-swagger
	swagger serve -F=swagger swagger.json