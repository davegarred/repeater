deps:
		go get -u github.com/google/uuid

clean:
		rm -f repeater

test:
		go test -v github.com/davegarred/repeater...

build: test
		go build -gcflags "-N -l"

run: build
		./repeater \
			-log /home/ubuntu/repeater.log \
			-disk

debug: build
		/home/ubuntu/go/go/bin/dlv debug
#		gdb /home/ubuntu/go/src/github.com/davegarred/repeater/repeater
