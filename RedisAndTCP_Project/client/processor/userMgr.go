package processor

import (
	"RedisAndTCP_Project/client/model"
	"RedisAndTCP_Project/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

//在用户登陆成功时，对Curuser初始化
var Curuser model.CurUser 

//在客户端显示当前在线用户
func outputOnlineUsers(){
	fmt.Println("当前用户列表：")
	for id,_ := range onlineUsers{
		fmt.Println("用户id:\t",id)
	}
}

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){

	user,ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok{
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	

	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUsers()
}
