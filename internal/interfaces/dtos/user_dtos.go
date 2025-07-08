package dtos

type LoginUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
