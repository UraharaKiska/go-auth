package model

import (
)


type UserAuth struct {
	Email string
	Password string
}

type UserBaseInfo struct {
	Email string `json:"email"`
	Role     string `json:"role"`
}

