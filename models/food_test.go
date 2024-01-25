package models

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	var foods []Food
	DB.Model(Food{}).Unscoped().Find(&foods)
	for _, food := range foods {
		fmt.Println(food.Name)

	}
}
