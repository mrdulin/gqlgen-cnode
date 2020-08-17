PROJECTNAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
MAIN_FILE=server.go
TEST=./...

build:
	@echo "  >  Building binary..."
	go build -o ${GOBIN}/${PROJECTNAME} ${MAIN_FILE}

model-gen:
	go run github.com/99designs/gqlgen generate
start:
	go run server.go
test-coverage:
	go test $$(go list $(TEST) | grep -v /mocks/ | grep -v /graph/) -v -short -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html
run: build
	${GOBIN}/${PROJECTNAME}