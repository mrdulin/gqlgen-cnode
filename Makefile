
model-gen:
	go run github.com/99designs/gqlgen generate
	go run modelGenHook.go 
start:
	go run server.go