package main

import (
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/rawsql"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../dal/gen",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	db, err := gorm.Open(rawsql.New(rawsql.Config{
		FilePath: []string{
			"item.sql",
			"recipe.sql",
		},
	}))
	if err != nil {
		panic(err)
	}
	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModelAs("recipes", "Recipe"),
		g.GenerateModelAs("recipe_items", "RecipeItem"),
		g.GenerateModelAs("items", "Item"),
	)
	// 执行并生成代码
	g.Execute()
}
