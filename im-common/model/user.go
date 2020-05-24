package model

import "go-im/im-common/enum"

type User struct {
	Id       int
	Account  string
	Password string
	Age      int
	Frineds  []FriendInfo
}

type FriendInfo struct {
	Account  string
	IsOnLine enum.EnumOnline
}
