NAME=IpProxyPool
BIN_DIR=bin
VERSION=$(shell git describe --tags || echo "unknown version")
GO_BUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-w -s'

PLATFORM_LIST = \
	darwin-amd64 \
	linux-386 \
	linux-amd64 \
	linux-armv5 \
	linux-armv6 \
	linux-armv7 \
	linux-armv8 \
	linux-mips-softfloat \
	linux-mips-hardfloat \
	linux-mipsle-softfloat \
	linux-mipsle-hardfloat \
	linux-mips64 \
	linux-mips64le \
	freebsd-386 \
	freebsd-amd64

all: linux-amd64 darwin-amd64

docker:
	$(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-386:
	GOARCH=386 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-armv5:
	GOARCH=arm GOOS=linux GOARM=5 $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-armv6:
	GOARCH=arm GOOS=linux GOARM=6 $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-armv7:
	GOARCH=arm GOOS=linux GOARM=7 $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-armv8:
	GOARCH=arm64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-mips-softfloat:
	GOARCH=mips GOMIPS=softfloat GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-mips-hardfloat:
	GOARCH=mips GOMIPS=hardfloat GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-mipsle-softfloat:
	GOARCH=mipsle GOMIPS=softfloat GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-mipsle-hardfloat:
	GOARCH=mipsle GOMIPS=hardfloat GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-mips64:
	GOARCH=mips64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-mips64le:
	GOARCH=mips64le GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

freebsd-386:
	GOARCH=386 GOOS=freebsd $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

freebsd-amd64:
	GOARCH=amd64 GOOS=freebsd $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

gz_releases=$(addsuffix .gz, $(PLATFORM_LIST))

$(gz_releases): %.gz : %
	chmod +x $(BIN_DIR)/$(NAME)-$(basename $@)
	gzip -f -S -$(VERSION).gz $(BIN_DIR)/$(NAME)-$(basename $@)

all-arch: $(PLATFORM_LIST)

releases: $(gz_releases)

clean:
	rm $(BIN_DIR)/*
