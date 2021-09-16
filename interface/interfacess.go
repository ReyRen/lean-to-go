package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB
	a = PhoneConnecter{name: "ppppppphoneConnecter"}
	b := PhoneConnecter{name: "xxxxxxConnecter"}
	a.Connect()
	DisConnect(b)

	pc := PhoneConnecter{
		name: "synonymy",
	}
	var c Connecter
	c = Connecter(pc)
	c.Connect() // no c.Name()

}

func DisConnect(usb interface{}) {
	/*if pc, ok := usb.(PhoneConnecter); ok {

		fmt.Println("DisConnect:", pc.name)
		return
	} else {
		fmt.Println("DisConnect unknown device")
	}*/

	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("DisConnect:", v.name)
	default:
		fmt.Println("DisConnect unknown device")
	}
}
