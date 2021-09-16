package utils

import "fmt"

type FamilyAccount struct {
	key     string  // store user input options
	loop    bool    // judge whether exit for loop
	balance float64 // account balance
	money   float64 // amount for each income/outcome
	note    string  // note of each income/outcome
	flag    bool    // record whether have income/outcome
	details string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "Income\tBalance\tAmount\tNotes",
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("----------------------LOG----------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("Nothing in logs...")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("income amount:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("income notes:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\nIncome\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) outcome() {
	fmt.Println("outcome amount:")
	fmt.Scanln(&this.money)

	if this.money > this.balance {
		fmt.Println("insufficient balance...")
	}
	this.balance -= this.money
	fmt.Println("outcome notes:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\nOutcome\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) exit() {
	fmt.Println("exit?y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("input again... y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n----------------------Family account APP----------------------")
		fmt.Println("                        1 income/outcome details")
		fmt.Println("                        2 income registry")
		fmt.Println("                        3 outcome registry")
		fmt.Println("                        4 exit")
		fmt.Print("(1--4)")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.outcome()
		case "4":
			this.exit()
		default:
			fmt.Println("input again...")
		}
		if !this.loop {
			break
		}
	}
}
