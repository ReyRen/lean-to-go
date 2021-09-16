package main

import (
	"fmt"
	"reflect"
)

type Users struct {
	Id   int
	Name string
	Age  int
}

func main() {
	/*x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)*/

	u := Users{
		Id:   99,
		Name: "xx",
		Age:  88,
	}
	Set(&u)

	/*var r io.Reader
	tty, _ := os.OpenFile("tty", os.O_RDONLY, 0)
	r = tty
	var w io.Writer
	w = r*/
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	//	fmt.Println(v.Kind(), v.Type())
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Bad")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("BYB")
	}

}
