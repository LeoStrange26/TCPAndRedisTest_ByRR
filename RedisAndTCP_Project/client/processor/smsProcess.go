package processor

import (
	"RedisAndTCP_Project/common/message"
	"RedisAndTCP_Project/common/utils"
	"encoding/json"
	"fmt"
)

type SmsProcessor struct {

}

//发送群聊消息
func (p *SmsProcessor) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = Curuser.UserId  
	smsMes.UserStatus = Curuser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化错误")
		return 
	}

	mes.Data = string(data)

	data1, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化错误")
		return
	}
	tf := &utils.Transfer{
		Conn : Curuser.Conn,
	}

	err = tf.WritePkg(data1)
	if err != nil{
		fmt.Println("写入消息错误")
		return
	}

	return
}