package model

type User struct {
	Id int32 `ddb:"id" json:"id"`
	Name string `ddb:"name" json:"name"`
	Password string `ddb:"password" json:"password"`
	Email string `ddb:"email" json:"email"`
}
