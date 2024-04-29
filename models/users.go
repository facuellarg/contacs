package models

type User struct {
	Id    int    `json:"id"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
