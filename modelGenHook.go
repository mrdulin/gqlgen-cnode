// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		if strings.Contains(strings.ToLower(model.Name), "params") {
			for _, field := range model.Fields {
				tagFieldName := strings.ReplaceAll(strings.Split(field.Tag, ":")[1], "\"", "")
				queryTag := fmt.Sprintf(`url:"%v,omitempty"`, tagFieldName)
				field.Tag = fmt.Sprintf(`%v %v`, field.Tag, queryTag)
			}
		}
	}
	return b
}

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
