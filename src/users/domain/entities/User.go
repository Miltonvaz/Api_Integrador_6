package entities

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Code     int    `json:"code"`
}
