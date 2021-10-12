package process

import (
	"encoding/json"
	"fmt"
	"lean-to-go/chatRoom-MVC/client/utils"
	"lean-to-go/chatRoom-MVC/common/message"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("---------------WELCOME XXX LOGIN SUCCESSFUL---------------")
	fmt.Println("\t\t\t 1 Display Online Users")
	fmt.Println("\t\t\t 2 Send Messages")
	fmt.Println("\t\t\t 3 Messages List")
	fmt.Println("\t\t\t 4 Exit System")
	fmt.Println("---------------PLEASE CHOOSE(1-4)---------------")

	var key int
	var content string

	smsProcess := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("Please Input Send Message:")
		fmt.Scanf("%s\n", &content)
		err := smsProcess.SendGroupMsg(content)
		fmt.Println("1111111111111111111111")
		if err != nil {
			fmt.Println("smsProcess.SendGroupMsg err=", err)
			return
		}
	case 3:
		fmt.Println("Messages List")
	case 4:
		fmt.Println("Exit System")
		os.Exit(0)
	default:
		fmt.Println("Input err......")
	}
}

//keep alive with server to receive from server
func ProcessServerMsg(conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg err=", err)
			return
		}
		fmt.Println("######Received From Server######")
		switch msg.Type {
		case message.NotifyUserStatusMsgType:
			var notifyUserStatusMsg message.NotifyUserStatusMsg
			err := json.Unmarshal([]byte(msg.Data), &notifyUserStatusMsg)
			if err != nil {
				fmt.Println("json.Unmarshal err=", err)
				return
			}
			updateUserStatus(&notifyUserStatusMsg)
			outputOnlineUser()
		case message.SmsMsgType:
			outputGroupMsg(&msg)
		default:
			fmt.Println("Unknown Message Type From Server.")
		}
		fmt.Println("######Received From Server######")
	}
}
