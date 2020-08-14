# gqlgen-cnode

A GraphQL APIs for [CNode](https://cnodejs.org) community built with [99designs/gqlgen](https://github.com/99designs/gqlgen) package.

### Usage

Every time you modify the GraphQL schema definition, you need to execute the `make model-gen` command to update the `root.resolver.go`, `model_gen.go` and `generated.go` files.
Then implement the corresponding resolvers. 

For this project, the specific business implementation is in the service layer.

More info, See [Makefile](./Makefile)