package process

import (
	"encoding/json"
	"fmt"
	"lean-to-go/chatRoom-MVC/client/utils"
	"lean-to-go/chatRoom-MVC/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMsg(content string) (err error) {
	var msg message.Message
	msg.Type = message.SmsMsgType

	// create SmsMsg
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = CurUser.UserId
	smsMsg.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMsg)
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
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err=", err)
		return
	}
	return
}
