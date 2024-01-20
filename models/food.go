package models

import (
	"select_menu/enums"
)

type Food struct {
	Base
	ID       uint             `gorm:"primaryKey"`
	Name     string           `gorm:"size:100;index:idx_name" json:"name"`
	Material string           `gorm:"size:255" json:"material"`
	Url      string           `gorm:"size:255" json:"url"`
	Status   enums.FoodStatus `gorm:"column:status;type:tinyint(1)" json:"status"` //1:热菜；2:凉菜；3:汤菜
}

func (table *Food) TableName() string {
	return "food"

}
