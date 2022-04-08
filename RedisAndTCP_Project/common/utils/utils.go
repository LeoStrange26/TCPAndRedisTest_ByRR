package utils

//工具
import (
	"RedisAndTCP_Project/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [4096]byte
}


//返回根据该连接读到的消息
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 4096)
	//读取前四位（消息长度）
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		// fmt.Println("read package head error")
		return
	}
	//pkgLen表示消息长度
	var pkgLen uint32 = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取内容
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if err != nil || n != int(pkgLen) {
		fmt.Println("read package body error")
		return
	}

	//反序列化消息
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("反序列化出错")
		return
	}
	return
}

//传入数据流，将数据流写入该连接
func (this *Transfer) WritePkg(data []byte) (err error) {
	//1、先发送一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	n, err := this.Conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("写入长度失败")
		return
	}
	//2、再写入消息
	n, err = this.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("写入数据流失败")
		return
	}
	return
}