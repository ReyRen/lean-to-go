package main

import (
	"fmt"
	"lean-to-go/customerManage/model"
	"lean-to-go/customerManage/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

// show all customer details
func (this *customerView) list() {
	customers := this.customerService.List()

	fmt.Println("----------------Customer List----------------")
	fmt.Println("Id\tName\t\tGender\tAge\tPhone\tEmail")
	for i := 0; i < len(customers); i++ {
		/*fmt.Println(customers[i].Id, "\t", customers[i].Name, "\t", customers[i].Gender, "\t", customers[i].Age,
		"\t", customers[i].Phone, "\t", customers[i].Email)*/
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n----------------Customer List End----------------\n\n")
}

func (this *customerView) Add() {
	fmt.Println("----------------Customer Add----------------")
	fmt.Println("Name:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("Gender:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("Age:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("Phone:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("Email:")
	email := ""
	fmt.Scanln(&email)

	//Id use system
	customer := model.NewCustomer2(name, gender, age, phone, email)
	this.customerService.Add(customer)
	fmt.Printf("\n----------------Customer Add End----------------\n\n")
}

func (this *customerView) DeleteById() {
	fmt.Println("----------------Customer Delete----------------")
	fmt.Println("Please choose Id:")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	if this.customerService.DeleteById(id) {
		fmt.Printf("\n----------------Customer Delete Done----------------\n\n")
	} else {
		fmt.Printf("\nId is Wrong...\n\n")

	}
}

//display main menu
func (this *customerView) mainMenu() {
	for {
		fmt.Println("----------------Customer Manager APP----------------")
		fmt.Println("                1 Add Customer")
		fmt.Println("                2 Modify Customer")
		fmt.Println("                3 Delete Customer")
		fmt.Println("                4 List Customer")
		fmt.Println("                5 Exit")
		fmt.Print("(1-5)")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.Add()
		case "2":
			fmt.Println("Modify Customer")
		case "3":
			this.DeleteById()
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("Wrong input...")
		}
		if !this.loop {
			break
		}
	}
}

func main() {
	customerview := customerView{
		key:  "",
		loop: true,
	}
	customerview.customerService = service.NewCustomerService()
	customerview.mainMenu()
}
