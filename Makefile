BINARY=omnia

VERSION=0.1
# BUILD=`date +%FT%T%z`
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

default: generate
	go build ${LDFLAGS} -o ${BINARY}

install: generate
	go install ${LDFLAGS}

generate:
	go generate ./...

updatedeps: # from line 2 is a temporary fix
	go get -u github.com/jteeuwen/go-bindata/...
	cd ${GOPATH}/src/github.com/jteeuwen/go-bindata ; \
		git remote add fork -f https://github.com/fridolin-koch/go-bindata.git 2>/dev/null; \
		git checkout -b fix_not_exist fork/master 2>/dev/null; \
		go install ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default install generate updatedeps clean
