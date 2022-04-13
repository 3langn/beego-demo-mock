package models

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Message string     `json:"message"`
	User    *User      `json:"user"`
	Token   *AuthToken `json:"token"`
}

type ResponseDto struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
