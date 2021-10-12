package process

import (
	"encoding/json"
	"fmt"
	"lean-to-go/chatRoom-MVC/client/utils"
	"lean-to-go/chatRoom-MVC/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}
	defer conn.Close()
	// conn send msg to server
	var msg message.Message
	msg.Type = message.LoginType
	// create LoginMsg struct
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd
	// LoginMsg serialize
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	msg.Data = string(data)

	// msg serialize
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	// data len send to server6
	// len to []byte convert
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg err=", err)
		return
	}

	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}

	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if loginResMsg.Code == 200 {

		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		fmt.Println("Online users:")
		for _, v := range loginResMsg.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("UserId:\t", v)

			// initialize
			onlineUsers[v] = &message.User{
				UserId:     v,
				UserPwd:    "",
				UserName:   "",
				UserStatus: message.UserOnline, // default
			}
		}

		fmt.Printf("\n\n")

		//goroutine: keep alive with server to receive from server
		go ProcessServerMsg(conn)

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMsg.Error)
	}

	return
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}
	defer conn.Close()

	// conn send msg to server
	var msg message.Message
	msg.Type = message.RegisterType
	// create LoginMsg struct
	var registerMsg message.RegisterMsg
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = userPwd
	registerMsg.User.UserName = userName
	// registerMsg serialize
	data, err := json.Marshal(&registerMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	msg.Data = string(data)

	// msg serialize
	data, err = json.Marshal(&msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	// data len send to server6
	// len to []byte convert
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg err=", err)
		return
	}

	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}
	var registerResMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if registerResMsg.Code == 200 {

		/*//goroutine: keep alive with server to receive from server
		go ProcessServerMsg(conn)
		for  {
			ShowMenu()
		}*/
		fmt.Println("Register success...")
	} else {
		fmt.Println(registerResMsg.Error)
	}
	os.Exit(0)
	return

}
