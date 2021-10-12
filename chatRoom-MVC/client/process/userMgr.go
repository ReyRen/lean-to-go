package process

import (
	"fmt"
	"lean-to-go/chatRoom-MVC/client/model"
	"lean-to-go/chatRoom-MVC/common/message"
)

var onlineUsers = make(map[int]*message.User, 10)
var CurUser model.CurUser

func outputOnlineUser() {
	fmt.Println("---------------Online Users:---------------")
	for id, _ := range onlineUsers {
		fmt.Println("UserID: \t", id)
	}
	fmt.Println("---------------Online Users:---------------")
}

func updateUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg) {

	user, ok := onlineUsers[notifyUserStatusMsg.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMsg.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMsg.Status

	onlineUsers[notifyUserStatusMsg.UserId] = user
}
