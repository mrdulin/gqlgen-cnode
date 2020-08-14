package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mrdulin/gqlgen-cnode/graph/generated"
	"github.com/mrdulin/gqlgen-cnode/graph/resolver"
	"github.com/mrdulin/gqlgen-cnode/services"
	appHttp "github.com/mrdulin/gqlgen-cnode/utils/http"
)

const defaultPort = "8080"
const BaseURL = "https://cnodejs.org/api/v1"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	baseUrl := os.Getenv("API_URL")
	if baseUrl == "" {
		baseUrl = BaseURL
	}
	hc := appHttp.NewClient()
	topicService := services.NewTopicService(hc, baseUrl)
	userService := services.NewUserService(hc, baseUrl)
	messageService := services.NewMessageService(hc, baseUrl)

	resolvers := resolver.Resolver{
		TopicService:   topicService,
		UserService:    userService,
		MessageService: messageService,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
