package model

import (
	"RedisAndTCP_Project/common/message"
	"net"
)

//当前在线用户
type CurUser struct {
	Conn net.Conn
	message.User
}