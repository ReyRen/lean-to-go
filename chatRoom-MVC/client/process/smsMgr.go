package process

import (
	"encoding/json"
	"fmt"
	"lean-to-go/chatRoom-MVC/common/message"
)

func outputGroupMsg(msg *message.Message) {

	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	info := fmt.Sprintf("UserId:\t%d say:\t%s", smsMsg.UserId, smsMsg.Content)
	fmt.Println(info)
	fmt.Println()
}
