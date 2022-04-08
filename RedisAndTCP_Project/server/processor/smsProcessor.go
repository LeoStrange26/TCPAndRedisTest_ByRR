package processor

import (
	"RedisAndTCP_Project/common/message"
	"RedisAndTCP_Project/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {

	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	//遍历服务器端的onlineUsers 并群发消息
	for id,up := range userMgr.onlineUsers{
		if id == smsMes.UserId{
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
	return
}

func (this *SmsProcess)SendMesToEachOnlineUser(data []byte,conn net.Conn) (err error) {
	//创建一个Transfer,发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("写入数据错误")
		return
	}
	return

}