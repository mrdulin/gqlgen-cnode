package main

import (
  "github.com/mrdulin/gqlgen-cnode/services"
  "github.com/mrdulin/gqlgen-cnode/utils"
  "log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mrdulin/gqlgen-cnode/graph/resolver"
	"github.com/mrdulin/gqlgen-cnode/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
  baseUrl := os.Getenv("API_URL")
  if baseUrl == "" {
    baseUrl = "https://cnodejs.org/api/v1"
  }
	httpClient := utils.NewHttpClient()
	topicService := services.NewTopicService(httpClient, baseUrl)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{TopicService: topicService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
