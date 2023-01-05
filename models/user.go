package models

type User struct {
	UserID   int64  `json:"user_id,string" gorm:"user_id;"`
	Username string `gorm:"username;not null"`
	Password string `gorm:"password;not null"`
	Token    string
}
