package gen

import (
	"github.com/richsoap/RecipeCalculator/dal/model"
)

func ItemToUpdateMap(i *model.Item) map[string]any {
	return map[string]any{
		Item.Name.ColumnName().String(): i.Name,
	}
}
