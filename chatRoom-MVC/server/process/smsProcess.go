package process

import (
	"encoding/json"
	"fmt"
	"lean-to-go/chatRoom-MVC/common/message"
	"lean-to-go/chatRoom-MVC/server/utils"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMsg(msg *message.Message) {

	//msg is already unmarshalled

	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marsha err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMsg.UserId {
			continue
		}
		this.SendMsgToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMsgToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err=", err)
	}
}
