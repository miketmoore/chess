GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=chess
REPO=github.com/miketmoore/chess
MAIN=cmd/chess/chess.go

all: deps test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)
test:
	$(GOTEST) -v
clean: 
	$(GOCLEAN) -i
	rm -f $(BINARY_NAME)
deps:
	$(GOGET) -u $(REPO)/...
