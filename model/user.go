package model

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment"`
	Password string `gorm:"size:255"`
	Name     string
	Email    string
	Age      int
	Phone    string
}
