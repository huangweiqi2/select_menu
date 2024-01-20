package enums

type FoodStatus int

const (
	FoodStatusHotDefault FoodStatus = iota
	FoodStatusHot
	FoodStatusCold
	FoodStatusSoup
)
