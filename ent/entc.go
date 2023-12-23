//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
	"github.com/tiagoposse/ogent"

	ogauth "github.com/tiagoposse/ogent-auth/extension"
)

func main() {
	spec := new(ogen.Spec)

	addProviderSpec := func(_ *gen.Graph, spec *ogen.Spec) error {
		spec.AddSchema("Provider", ogen.NewSchema().AddRequiredProperties(ogen.String().ToProperty("ApiKey"), ogen.String().ToProperty("Endpoint")))
		return nil
	}

	authzExt := ogauth.NewOgentAuthExtension(
		ogauth.WithApiKeySecurity(),
		ogauth.WithCookieSecurity(),
		ogauth.WithDefaultGlobalAuthentication(),
	)
	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(
			addCustomPaths,
			authzExt.SecurityMutation(),
			addProviderSpec,
		),
	)

	// tf := entform.NewTerraformExtension("Launchpad")
	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureNamedEdges,
			gen.FeatureUpsert,
			gen.FeatureIntercept,
			gen.FeaturePrivacy,
		},
	}, entc.Extensions(ogent, authzExt, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func addCustomPaths(graph *gen.Graph, spec *ogen.Spec) error {
	// spec.Paths["/accesses/{id}/approvals"].Get.Tags = append(spec.Paths["/accesses/{id}/approvals"].Get.Tags, "x-tf-ignore")
	spec.Paths["/accesses/{id}/access-tokens"].Get.Tags = append(spec.Paths["/accesses/{id}/access-tokens"].Get.Tags, "x-tf-ignore")

	spec.Paths["/audits"].Get.Tags = append(spec.Paths["/audits"].Get.Tags, "x-tf-ignore")
	spec.Paths["/audits/{id}"].Get.Tags = append(spec.Paths["/audits/{id}"].Get.Tags, "x-tf-ignore")

	spec.Paths["/approvals"].Get.Tags = append(spec.Paths["/approvals"].Get.Tags, "x-tf-ignore")
	spec.Paths["/approvals"].Post.Tags = append(spec.Paths["/approvals"].Post.Tags, "x-tf-ignore")
	spec.Paths["/approvals/{id}"].Get.Tags = append(spec.Paths["/approvals/{id}"].Get.Tags, "x-tf-ignore")
	spec.Paths["/approvals/{id}"].Patch.Tags = append(spec.Paths["/approvals/{id}"].Patch.Tags, "x-tf-ignore")
	spec.Paths["/approvals/{id}"].Delete.Tags = append(spec.Paths["/approvals/{id}"].Delete.Tags, "x-tf-ignore")

	return nil
}
