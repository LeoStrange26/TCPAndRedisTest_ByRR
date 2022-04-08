package processor

import (
	"RedisAndTCP_Project/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化出错")
		return
	}
	info := fmt.Sprintf("用户%d\t发送了消息：\t%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}