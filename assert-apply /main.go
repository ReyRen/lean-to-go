package main

import "fmt"

type Usb interface {
	Plugin()
	Plugout()
}

type Phone struct {
	name string
}

func (p Phone) Plugin() {
	fmt.Printf("phone %s plugged in...\n", p.name)
}
func (p Phone) Plugout() {
	fmt.Printf("phone %s plugged out...\n", p.name)
}
func (p Phone) Call() {
	fmt.Printf("phone %s is calling...\n", p.name)
}

type Camera struct {
	name string
}

func (c Camera) Plugin() {
	fmt.Printf("camera %s plugged in...\n", c.name)
}
func (c Camera) Plugout() {
	fmt.Printf("camera %s plugged out...\n", c.name)
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Plugin()
	//assert
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Plugout()
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{name: "vivo"}
	usbArr[1] = Phone{name: "iphone"}
	usbArr[2] = Camera{name: "niko"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
