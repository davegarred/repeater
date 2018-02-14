deps:
		go get -u github.com/google/uuid
		go get -u gopkg.in/h2non/filetype.v1
		go get -u google.golang.org/grpc
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go

proto:
		protoc -I grpcfile --go_out=plugins=grpc:grpcfile grpcfile/*.proto

clean:
		go clean
		rm -f repeater
		rm -f grpcclient/grpcclient

test:
		go test -cover -v github.com/davegarred/repeater...

build: test proto
		go build -gcflags "-N -l"

run: build
		./repeater \
			-log /home/ubuntu/repeater.log \
			-disk

debug: build
		/home/ubuntu/go/go/bin/dlv debug
#		gdb /home/ubuntu/go/src/github.com/davegarred/repeater/repeater
