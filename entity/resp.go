package entity

type RandomResp struct {
	Foods     []FoodResp `json:"foods"`
	Materials []string   `json:"materials"` //总配料
}
type FoodResp struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Material []string `json:"material"`
	Status   int      `json:"status"`
}
