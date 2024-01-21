package enums

type FoodStatus int

const (
	FoodStatusHotDefault FoodStatus = iota
	FoodStatusHot
	FoodStatusCold
	FoodStatusSoup
)

func (f FoodStatus) IsHot() bool {
	return f == FoodStatusHot
}
func (f FoodStatus) IsCold() bool {
	return f == FoodStatusCold
}

func (f FoodStatus) IsSoup() bool {
	return f == FoodStatusSoup
}

func (f FoodStatus) Int() int {
	return int(f)
}
