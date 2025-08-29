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

/*
package api
type Item struct {
	ID        int64  `thrift:"id,1" form:"id" json:"id" query:"id"`
	Name      string `thrift:"name,2" form:"name" json:"name" query:"name"`
	CreatedAt string `thrift:"created_at,3" form:"created_at" json:"created_at" query:"created_at"`
	UpdatedAt string `thrift:"updated_at,4" form:"updated_at" json:"updated_at" query:"updated_at"`
}

package model
type Item struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

*/
