package models

type User struct {
	Base
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"column:username;type:varchar(100)" json:"username"`
	Password string `gorm:"column:password;type:varchar(100)" json:"password"`
	Email    string `gorm:"column:;email;type:varchar(100)" json:"email"`
	Phone    string `gorm:"column:phone;type:varchar(20)" json:"phone"`
}

func (table *User) TableName() string {
	return "user"

}
