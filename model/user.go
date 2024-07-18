package model

type UserRegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PasswordLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserListReq struct {
	Page int `json:"page" query:"page"`
	Rows int `json:"rows" query:"rows"`
}

type UserCreateReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
