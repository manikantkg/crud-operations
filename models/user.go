package models

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}
