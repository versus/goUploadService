GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
EXE  := upload
PKG  := github.com/versus/gouploadservice
VER := 0.0.2

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
