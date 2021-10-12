package main

import (
	"fmt"
	"lean-to-go/chatRoom-MVC/client/process"
)

var (
	userId   int
	userPwd  string
	userName string
)

func main() {

	var key int

	for {
		fmt.Println("---------------WELCOME LOGIN CHAT ROOT---------------")
		fmt.Println("\t\t\t 1 Login")
		fmt.Println("\t\t\t 2 Registry")
		fmt.Println("\t\t\t 3 Exit")
		fmt.Println("---------------PLEASE CHOOSE(1-3)---------------")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("Please input user ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("Please input user PWD:")
			fmt.Scanf("%s\n", &userPwd)

			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("up.Login err=", err)
				return
			}
		case 2:
			fmt.Println("Please input user ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("Please input user PWD:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("Please input user NAME:")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			err := up.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("up.Register err=", err)
				return
			}
		case 3:
			fmt.Println("exit")
		default:
			fmt.Println("Input err......")
		}
	}
}
