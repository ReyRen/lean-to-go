package process

import (
	"encoding/json"
	"fmt"
	"lean-to-go/chatRoom-MVC/common/message"
	"lean-to-go/chatRoom-MVC/server/model"
	"lean-to-go/chatRoom-MVC/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}
func (this *UserProcess) NotifyMeOnline(userId int) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType

	var notifyUserStatusMsg message.NotifyUserStatusMsg
	notifyUserStatusMsg.UserId = userId
	notifyUserStatusMsg.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err=", err)
		return
	}
}

func (this *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	//1. get data from msg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType
	var logResMsg message.LoginResMsg

	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			logResMsg.Code = 500
			logResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			logResMsg.Code = 300
			logResMsg.Error = err.Error()
		} else {
			logResMsg.Code = 404
			logResMsg.Error = "???? Error"
		}
	} else {
		logResMsg.Code = 200
		this.UserId = loginMsg.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(loginMsg.UserId)
		for id, _ := range userMgr.onlineUsers {
			logResMsg.UserIds = append(logResMsg.UserIds, id)
		}
		fmt.Println(user, "login success")
	}

	data, err := json.Marshal(logResMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(msg *message.Message) (err error) {
	//1. get data from msg
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.RegisterResType
	var registerResMsg message.RegisterResMsg

	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMsg.Code = 505
			registerResMsg.Error = err.Error()
		} else {
			registerResMsg.Code = 506
			registerResMsg.Error = "???? register 506"
		}
	} else {
		registerResMsg.Code = 200
	}
	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
