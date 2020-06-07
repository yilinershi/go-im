package model

import (
	"go-im/im-common/enum"
	"go-im/im-common/zinx/iface"
)

type User struct {
	Id       int
	Account  string
	Password string
	Age      int
	Frineds  []FriendInfo
	Conn     iface.IConnection
}

type FriendInfo struct {
	Account  string
	IsOnLine enum.EnumOnline
}
