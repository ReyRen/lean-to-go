package main

import (
	"errors"
	"fmt"
	"io"
	"lean-to-go/chatRoom-MVC/common/message"
	"lean-to-go/chatRoom-MVC/server/process"
	"lean-to-go/chatRoom-MVC/server/utils"
	"net"
)

type Processer struct {
	Conn net.Conn
}

func (this *Processer) serverProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginType:
		// login logicala
		up := &process.UserProcess{Conn: this.Conn}
		err = up.ServerProcessLogin(msg)
	case message.RegisterType:
		// register logical
		up := &process.UserProcess{Conn: this.Conn}
		err = up.ServerProcessRegister(msg)
	case message.SmsMsgType:
		// create a SmsProcess
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMsg(msg)
	default:
		fmt.Println("Message Type not exist...")
	}
	return
}

func (this *Processer) processMainWorker() (err error) {

	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client is loggout, server connection also exit...")
				return err
			} else {
				err = errors.New("readPkg error")
				return err
			}
		}
		err = this.serverProcessMsg(&msg)
		if err != nil {
			fmt.Println("serverProcessMsg err=", err)
			return err
		}
	}
}
