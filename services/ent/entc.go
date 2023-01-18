//go:build ignore

package main

import (
	"entgo.io/ent/entc/gen"
	"log"

	"entgo.io/ent/entc"
	"github.com/hedwigz/entviz"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{Features: []gen.Feature{gen.FeatureUpsert, gen.FeatureVersionedMigration}}, entc.Extensions(entviz.Extension{})); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
