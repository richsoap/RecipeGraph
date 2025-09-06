package convert

import (
	"time"

	"github.com/bytedance/gg/gptr"
	api "github.com/richsoap/RecipeCalculator/biz/model/recipe/api"
	model "github.com/richsoap/RecipeCalculator/dal/model"
)

// ConvertRecipeItemToApi 将数据库模型转换为API模型
func ConvertRecipeItemToApi(recipeItem *model.RecipeItem, item *model.Item) *api.RecipeItem {
	if recipeItem == nil {
		return nil
	}

	return &api.RecipeItem{
		ID:        gptr.Of(int64(recipeItem.ID)),
		RecipeID:  gptr.Of(int64(recipeItem.RecipeID)),
		Space:     gptr.Of(int64(recipeItem.Space)),
		Type:      gptr.Of(api.RecipeItemType(recipeItem.Type)),
		Item:      ConvertItemToApi(item),
		Count:     &recipeItem.Count_,
		CreatedAt: gptr.Of(recipeItem.CreatedAt.Format(time.RFC3339)),
		UpdatedAt: gptr.Of(recipeItem.UpdatedAt.Format(time.RFC3339)),
	}
}

// ConvertRecipeItemToModel 将API模型转换为数据库模型
func ConvertRecipeItemToModel(recipeItem *api.RecipeItem) (*model.RecipeItem, error) {
	if recipeItem == nil {
		return nil, nil
	}

	createdAt, err := time.Parse(time.RFC3339, recipeItem.GetCreatedAt())
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, recipeItem.GetUpdatedAt())
	if err != nil {
		return nil, err
	}

	return &model.RecipeItem{
		ID:        int32(recipeItem.GetID()),
		RecipeID:  int32(recipeItem.GetRecipeID()),
		Space:     int32(recipeItem.GetSpace()),
		Type:      int32(recipeItem.GetType()),
		ItemID:    int32(recipeItem.GetItem().GetID()),
		Count_:    recipeItem.GetCount(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

/*
api
type RecipeItem struct {
	ID        int64          `thrift:"id,1,required" form:"id,required" json:"id,required" query:"id,required"`
	RecipeID  int64          `thrift:"recipe_id,2,required" form:"recipe_id,required" json:"recipe_id,required" query:"recipe_id,required"`
	Space     int64          `thrift:"space,3,required" form:"space,required" json:"space,required" query:"space,required"`
	Type      RecipeItemType `thrift:"type,4,required" form:"type,required" json:"type,required" query:"type,required"`
	ItemID    int64          `thrift:"item_id,5,required" form:"item_id,required" json:"item_id,required" query:"item_id,required"`
	Count     int32          `thrift:"count,6,required" form:"count,required" json:"count,required" query:"count,required"`
	CreatedAt string         `thrift:"created_at,7,required" form:"created_at,required" json:"created_at,required" query:"created_at,required"`
	UpdatedAt string         `thrift:"updated_at,8,required" form:"updated_at,required" json:"updated_at,required" query:"updated_at,required"`
}


model
type RecipeItem struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RecipeID  int32     `gorm:"column:recipe_id;not null" json:"recipe_id"`
	Space     int32     `gorm:"column:space;not null" json:"space"`
	Type      int32     `gorm:"column:type;not null" json:"type"`
	ItemID    int32     `gorm:"column:item_id;not null" json:"item_id"`
	Count_    int32     `gorm:"column:count;not null" json:"count"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

*/
