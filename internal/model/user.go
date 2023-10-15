package model

import "time"

type User struct {
	Id       int64
	Email    string
	Password string

	Birthday  string
	NickName  string
	Introduce string

	CTime time.Time
}

type SendUserProfile struct {
	Email string

	Birthday  string
	NickName  string
	Introduce string
}
