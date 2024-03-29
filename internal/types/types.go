// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

type JwtTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpireAt    int    `json:"expire_at"`
}

type RegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
}

type RandomRequest struct {
	HotNum  int `form:"hot_num"`
	ColdNum int `form:"cold_num"`
	SoupNum int `form:"soup_num"`
}

type RandomResponse struct {
	Foods     []FoodResponse `json:"foods"`
	Materials []string       `json:"materials"` //总配料
}

type FoodResponse struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Material []string `json:"material"`
	Status   int      `json:"status"`
}

type GetByNameRequest struct {
	Name string `form:"name"`
}

type GetByMaterialRequest struct {
	Material string `form:"material"`
}
