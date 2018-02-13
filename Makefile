deps:
		go get -u github.com/google/uuid

proto:
		protoc -I grpc/proto --go_out=plugins=grpc:grpc/proto grpc/proto/*.proto

clean:
		rm -f repeater

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
