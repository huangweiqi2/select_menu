package models

import (
	"select_menu/enums"
	"select_menu/internal/types"
	"strings"
)

type Food struct {
	Base
	Name     string           `gorm:"size:100;index:idx_name" json:"name"`
	Material string           `gorm:"size:255" json:"material"`
	Url      string           `gorm:"size:255" json:"url"`
	Status   enums.FoodStatus `gorm:"column:status;type:tinyint(1)" json:"status"` //1:热菜；2:凉菜；3:汤菜
}

func (f Food) TableName() string {
	return "food"
}
func (f Food) Response() types.FoodResponse {
	return types.FoodResponse{
		ID:       f.ID,
		Name:     f.Name,
		Material: strings.Split(strings.TrimPrefix(f.Material, "原料："), "、"),
		Status:   f.Status.Int(),
	}
}
