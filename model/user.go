package model

var AllowUser = map[string]string{}

func init() {
	AllowUser["admin"] = "admin"
	AllowUser["rahul"] = "rahul"
	AllowUser["test"] = "test"
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // for testing purpose storing
}

type UserResponse struct {
	Id      string      `json:"id,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
