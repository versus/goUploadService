GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
EXE  := upload
PKG  := github.com/versus/gouploadservice
VER := 0.0.2
current_dir := $(shell pwd)

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run:
	go run main.go

.PHONY:darwin linux 
darwin linux:
	GOOS=$@ CGO_ENABLED=0 GO111MODULE=on go build -o ./dist/$(EXE)-$(VER)-$@-$(GOARCH) $(PKG)

.PHONY: clean
clean:
	rm -rf ./dist/

.PHONY: docker-build
docker-build:
	docker build -f Dockerfile -t versus/go-upload:latest .

.PHONY: docker-run
docker-run:
	docker-compose -f docker-compose.yml up --build

.PHONY: docker-test
docker-test:
	cd terratest && go test -v -timeout 30m
	pwd && docker run --rm  -v $(current_dir)/goss.yaml:/goss.yaml -v /var/run/docker.sock:/var/run/docker.sock -e GOSS_FILES_STRATEGY=cp versus197/dgoss-docker-image:latest /usr/local/bin/dgoss run --entrypoint=/dgoss/start.sh versus/go-upload:test

	
