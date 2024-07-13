package model

import "time"

type User struct {
	Id        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	Email     string    `json:"email" db:"email"`
	Admin     int       `json:"admin" db:"admin"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

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
