package model

import (
	"lean-to-go/chatRoom-MVC/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
