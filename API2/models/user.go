package models

type User struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
