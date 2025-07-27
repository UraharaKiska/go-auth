package model

import (
	"database/sql"
	"time"
)

type OptionString struct {
	Value string
	Valid bool
}

type User struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
}

type UserUpdateInfo struct {
	Name  OptionString
	Email OptionString
}
