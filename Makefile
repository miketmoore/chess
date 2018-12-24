GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=chess

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v
clean: 
	$(GOCLEAN) -i
	rm -f $(BINARY_NAME)
deps:
	$(GOGET) github.com/faiface/pixel
	$(GOGET) github.com/BurntSushi/toml
	$(GOGET) github.com/nicksnyder/go-i18n/v2/i18n
