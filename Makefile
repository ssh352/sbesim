GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
DEPCMD=dep
BINARY=sbesim

all: install build

build:
	$(GOBUILD) -v -o $(BINARY)

run:
	./$(BINARY)

.PHONY : install
install:
	$(DEPCMD) ensure

.PHONY : clean
clean:
	rm -f $(BINARY) && rm -rf vendor && rm -f inttest

.PHONY : fmt
fmt:
	$(GOCMD) fmt ./

.PHONY : mockgen
mockgen:
	rm -rf ./mock/* && $(GOCMD) generate ./...

.PHONY : test
test:
	$(GOCMD)  test sbe/... -cover

.PHONY : inttest
inttest:
	cd integration && $(GOBUILD) -v -o inttest && mv inttest ../ && cd .. && ./inttest

.PHONY: help

# Show this help.
help:
	@echo all -- install dependencies and build binary
	@echo build -- build binary
	@echo install -- install dependencies
	@echo clean -- clean up dependencies and binaries
	@echo fmt -- auto format file
	@echo mockgen -- generate mock package
	@echo test -- run unit test
	@echo inttest -- run integration test
	@echo run -- start service
