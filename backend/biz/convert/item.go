package convert

import (
	"time"

	api "github.com/richsoap/RecipeCalculator/biz/model/recipe/api"
	model "github.com/richsoap/RecipeCalculator/dal/model"
)

// ConvertItemToApi 将数据库模型转换为API模型
func ConvertItemToApi(item *model.Item) *api.Item {
	if item == nil {
		return nil
	}

	return &api.Item{
		ID:        int64(item.ID),
		Name:      item.Name,
		CreatedAt: item.CreatedAt.Format(time.RFC3339),
		UpdatedAt: item.UpdatedAt.Format(time.RFC3339),
	}
}

// ConvertItemToModel 将API模型转换为数据库模型
func ConvertItemToModel(item *api.Item) (*model.Item, error) {
	if item == nil {
		return nil, nil
	}

	createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &model.Item{
		ID:        int32(item.ID),
		Name:      item.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
