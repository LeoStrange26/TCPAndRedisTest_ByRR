package processor

import (
	"RedisAndTCP_Project/common/message"
	"RedisAndTCP_Project/common/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登陆成功界面
func ShowMenu() {
	fmt.Println("\t----------登录成功，欢迎回来！----------")
	fmt.Println("\t\t  1、显示在线用户列表")
	fmt.Println("\t\t  2、发送消息")
	fmt.Println("\t\t  3、信息列表")
	fmt.Println("\t\t  4、退出系统")
	fmt.Println("\t\t   请选择(1-4)")
	fmt.Println("\t--------------------------------------")
	var key int
	var content string
	fmt.Scanf("%d\n",&key)
	//因为总会使用SmsProcessor实例（发送消息）,因此将其创建在switch外(虽然并没卵用)
	smsProcessor := &SmsProcessor{}
	switch key {
		case 1:
			fmt.Println("显示在线用户列表")
			outputOnlineUsers()
			
		case 2:
			// fmt.Println("发送消息")
			fmt.Println("请输入你要发送的消息：")
			fmt.Scanf("%s\n",&content)
			smsProcessor.SendGroupMes(content)
		case 3:
			fmt.Println("信息列表")

		case 4:
			fmt.Println("你选择了退出系统")
			os.Exit(0)

		default://其他非法输入
			fmt.Println("你的输入有误,请重新输入")

	}
}


func serverProcessMes(conn net.Conn){
	tf := &utils.Transfer{
		Conn : conn,
	}
	for{
		mes, err := tf.ReadPkg()
		if err != nil{
			fmt.Println("tf.ReadPkg error = ", err)
		}
		// fmt.Printf("message = %v\n",mes)
		switch mes.Type{
			case message.NotifyUserStatusMesType:
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				updateUserStatus(&notifyUserStatusMes)
			case message.SmsMesType:
				outputGroupMes(&mes)
			default:
				fmt.Println("服务器端返回了未知类型的消息")
		}
	}
}