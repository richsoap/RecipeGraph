package gen

import "github.com/richsoap/RecipeCalculator/dal/model"

func RecipeToUpdateMap(i *model.Recipe) map[string]any {
	return map[string]any{
		Recipe.ItemID.ColumnName().String():     i.ItemID,
		Recipe.Efficiency.ColumnName().String(): i.Efficiency,
	}
}
