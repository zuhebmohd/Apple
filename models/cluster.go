package models

type Cluster struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	ServerCount int    `json:"server_count"`
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
