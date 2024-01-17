package models

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Material string `gorm:"column:name;type:varchar(255)" json:"material"`
	Status   int    `gorm:"column:status;type:tinyint(1)" json:"status"`
}

func (table *Food) TableName() string {
	return "food"

}
