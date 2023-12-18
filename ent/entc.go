//go:build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

func main() {
	spec := new(ogen.Spec)

	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(
			addCustomPaths,
		),
	)

	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureNamedEdges,
			gen.FeatureUpsert,
			gen.FeatureIntercept,
		},
	}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func addCustomPaths(graph *gen.Graph, spec *ogen.Spec) error {
	// authHeader := map[string]*ogen.Parameter{
	// 	"Set-Cookie": {Schema: ogen.String(), Required: true},
	// }

	// locationHeader := map[string]*ogen.Parameter{
	// 	"Location": {Schema: &ogen.Schema{Type: "string", Format: "uri"}, Required: true},
	// }

	// samlAuthHeader := map[string]*ogen.Parameter{
	// 	"Set-Cookie": {Schema: ogen.String(), Required: true},
	// 	"Location":   {Schema: &ogen.Schema{Type: "string", Format: "uri"}, Required: true},
	// }

	return nil
}
