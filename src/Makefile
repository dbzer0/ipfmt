BINARY=ip2fmt

VERSION=0.0.1
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

all	: build

build:
	go build ${LDFLAGS} -o ../bin/${BINARY}

clean:
	if [ -f ../bin/${BINARY} ] ; then rm ../bin/${BINARY} ; fi

.PHONY: clean build
