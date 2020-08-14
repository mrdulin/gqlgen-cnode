package main

import (
	"log"
	"net/http"
	"os"

	httpClient "github.com/mrdulin/gqlgen-cnode/utils/httpClient"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mrdulin/gqlgen-cnode/graph/generated"
	"github.com/mrdulin/gqlgen-cnode/graph/resolver"
	"github.com/mrdulin/gqlgen-cnode/services"
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
	hc := httpClient.New()
	topicService := services.NewTopicService(hc, baseUrl)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{TopicService: topicService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
