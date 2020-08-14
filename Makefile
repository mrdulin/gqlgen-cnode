PROJECTNAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
MAIN_FILE=server.go

build:
	@echo "  >  Building binary..."
	go build -o ${GOBIN}/${PROJECTNAME} ${MAIN_FILE}

model-gen:
	go run github.com/99designs/gqlgen generate
	go run modelGenHook.go 
start:
	go run server.go
run: build
	${GOBIN}/${PROJECTNAME}