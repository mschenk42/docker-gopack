VERSION=v0.0.1
BINARY=docker-runpack

LDFLAGS=-ldflags "-X main.Version=${VERSION}"
GOOS=linux
GOARCH=amd64
FLAGS_linux_amd64= GOOS=${GOOS} GOARCH=${GOARCH}
ARCHIVE=${BINARY}-v${VERSION}-${GOARCH}-${GOOS}.tar.gz

build:
	$(FLAGS_linux_amd64) go build ${LDFLAGS} -o ./bin/${BINARY}
	# shasum -a 256 ./bin/${BINARY}
	# go build ${LDFLAGS} -o ${BINARY}

tar:
	tar -czvf ./bin/${ARCHIVE} ./bin/${BINARY}

test:
	go test $$(go list ./... | grep -v vendor)

vet:
	go vet $$(go list ./... | grep -v vendor)

clean:
	rm ./bin/${BINARY} ./bin/${ARCHIVE}

.PHONY: test vet install clean fmt tar
