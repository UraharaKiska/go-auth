package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	Name            string `db:"name"`
	Email           string `db:"email"`
	PasswordHash    string `db:"password_hash"`
	Role            string `db:"role"`
}

type UpdateUserInfo struct {
	Name  sql.NullString
	Email sql.NullString
}
