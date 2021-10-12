package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"lean-to-go/chatRoom-MVC/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (msg message.Message, err error) {

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		err = errors.New("read pkg header error")
		return
	}
	//fmt.Println("readed buf is: ", buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg data error")
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &msg)
	if err != nil {
		err = errors.New("unmarshal msg struct error")
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	// len
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	return
}
