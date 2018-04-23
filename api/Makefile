proto:
	protoc -I../idl \
		-I/usr/local/include \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		../idl/**.proto
	protoc -I../idl \
		-I/usr/local/include \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		../idl/**.proto

build: 
	go build -o /tmp/lifemgr-api github.com/smonv/lifemgr/api/cmd/api

install: build
	mv /tmp/lifemgr-api ${GOBIN}