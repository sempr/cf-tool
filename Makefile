NAME=cf
BINDIR=bin
VER := $(shell git rev-parse HEAD)
DT := $(shell date +%Y%m%d%H%M%S.%Z)
GOBUILD=go build -ldflags="-X 'main.Version=${DT}_$(VER)'" 

all: prepare linux-386 linux-amd64 linux-armv8 windows-386 windows-amd64 windows-arm64 darwin-amd64 darwin-arm64

prepare:
	mkdir -p $(BINDIR)

darwin-amd64:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

darwin-arm64:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

windows-386:
	CGO_ENABLED=0 GOARCH=386 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

windows-amd64:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

windows-arm64:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

linux-386:
	CGO_ENABLED=0 GOARCH=386 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-amd64:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-armv8:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@


