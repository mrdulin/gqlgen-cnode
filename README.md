# gqlgen-cnode

[![Go Report Card](https://goreportcard.com/badge/github.com/mrdulin/gqlgen-cnode)](https://goreportcard.com/report/github.com/mrdulin/gqlgen-cnode)

A GraphQL APIs for [CNode](https://cnodejs.org) community built with [99designs/gqlgen](https://github.com/99designs/gqlgen) package.

### Usage

If you created go model struct manually and split these model into different files rather than the model file specified in `gqlgen.yml` file. 
The struct tags should be updated manually too. Because these model files will NOT override by running `go run github.com/99designs/gqlgen generate` command.

If you created and updated go model struct using `modelgen` of `gqlgen` and don't split these models to different files.
If you want to add custom struct tags to go model struct, you need to run custom [modelgen hook]('./utils/hook/modelGen.go') after running `go run github.com/99designs/gqlgen generate` command.
This hook will update the struct tags to the generated models. 

Every time you modify the GraphQL schema definition, you need to execute the `make model-gen` command to update the `root.resolver.go`, `model_gen.go` and `generated.go` files.
Then implement the corresponding resolvers. 

For this project, the specific business implementation is in the service layer.

More info, See [Makefile](./Makefile)