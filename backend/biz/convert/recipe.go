package convert

import (
	"time"

	api "github.com/richsoap/RecipeCalculator/biz/model/recipe/api"
	model "github.com/richsoap/RecipeCalculator/dal/model"
)

// ConvertRecipeToApi 将数据库模型转换为API模型
func ConvertRecipeToApi(recipe *model.Recipe) *api.Recipe {
	if recipe == nil {
		return nil
	}

	return &api.Recipe{
		ID:         int64(recipe.ID),
		Space:      int64(recipe.Space),
		ItemID:     int64(recipe.ItemID),
		Efficiency: recipe.Efficiency,
		CreatedAt:  recipe.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  recipe.UpdatedAt.Format(time.RFC3339),
	}
}

// ConvertRecipeToModel 将API模型转换为数据库模型
func ConvertRecipeToModel(recipe *api.Recipe) (*model.Recipe, error) {
	if recipe == nil {
		return nil, nil
	}

	createdAt, err := time.Parse(time.RFC3339, recipe.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, recipe.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &model.Recipe{
		ID:         int32(recipe.ID),
		Space:      int32(recipe.Space),
		ItemID:     int32(recipe.ItemID),
		Efficiency: recipe.Efficiency,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}, nil
}

/*

api
// Recipe related structures
type Recipe struct {
	ID         int64   `thrift:"id,1,required" form:"id,required" json:"id,required" query:"id,required"`
	Space      int64   `thrift:"space,2,required" form:"space,required" json:"space,required" query:"space,required"`
	ItemID     int64   `thrift:"item_id,3,required" form:"item_id,required" json:"item_id,required" query:"item_id,required"`
	Efficiency float64 `thrift:"efficiency,4,required" form:"efficiency,required" json:"efficiency,required" query:"efficiency,required"`
	CreatedAt  string  `thrift:"created_at,5,required" form:"created_at,required" json:"created_at,required" query:"created_at,required"`
	UpdatedAt  string  `thrift:"updated_at,6,required" form:"updated_at,required" json:"updated_at,required" query:"updated_at,required"`
}

model
type Recipe struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Space      int32     `gorm:"column:space;not null" json:"space"`
	ItemID     int32     `gorm:"column:item_id;not null" json:"item_id"`
	Efficiency float64   `gorm:"column:efficiency;not null;default:1.0" json:"efficiency"`
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

*/
