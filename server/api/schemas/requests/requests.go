package requests

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdmin bool   `json:"isAdmin"`
}
