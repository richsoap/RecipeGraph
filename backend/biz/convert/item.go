package convert

import (
	"time"

	"github.com/bytedance/gg/gptr"
	api "github.com/richsoap/RecipeCalculator/biz/model/recipe/api"
	model "github.com/richsoap/RecipeCalculator/dal/model"
)

// ConvertItemToApi 将数据库模型转换为API模型
func ConvertItemToApi(item *model.Item) *api.Item {
	if item == nil {
		return nil
	}

	return &api.Item{
		ID:        gptr.Of(int64(item.ID)),
		Name:      gptr.Of(item.Name),
		CreatedAt: gptr.Of(item.CreatedAt.Format(time.RFC3339)),
		UpdatedAt: gptr.Of(item.UpdatedAt.Format(time.RFC3339)),
	}
}

// ConvertItemToModel 将API模型转换为数据库模型
func ConvertItemToModel(item *api.Item) (*model.Item, error) {
	if item == nil {
		return nil, nil
	}

	createdAt, err := time.Parse(time.RFC3339, item.GetCreatedAt())
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, item.GetUpdatedAt())
	if err != nil {
		return nil, err
	}

	return &model.Item{
		ID:        item.GetID(),
		Name:      item.GetName(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
