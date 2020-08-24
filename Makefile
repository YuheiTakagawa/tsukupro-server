.PHONY: all
.PHONY: run
.PHONY: pb

all: pb run


run: pb main.go
	go run main.go

pb: user.proto proto
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:./proto *.proto
