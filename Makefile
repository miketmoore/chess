GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=chess

all: deps test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v
clean: 
	$(GOCLEAN) -i
	rm -f $(BINARY_NAME)
deps:
	$(GOGET) -u github.com/miketmoore/chess/...
