package models

import (
	"errors"
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
func TestPanic(t *testing.T) {

	err := sdfdsfds()
	if err != nil {
		fmt.Println("sdfdsfds err:", err.Error())
	}
}

func sdfdsfds() error {
	return errors.New("wosile")
}
