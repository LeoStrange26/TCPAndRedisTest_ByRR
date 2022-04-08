package processor

import (
	"RedisAndTCP_Project/common/message"
	"RedisAndTCP_Project/common/utils"
	"RedisAndTCP_Project/server/model"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct{
	Conn net.Conn
	UserId int
}

//处理登陆
func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil{
		fmt.Println("反序列化失败 err = ",err)
		return 
	}

	//声明一个Message
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//再声明一个LoginResMes
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	if err != nil {
		switch err {
			case model.ERROR_USER_NOTEXIST:
				loginResMes.Code = 500
				loginResMes.Error = err.Error()
			case model.ERROR_USER_PWD:
				loginResMes.Code = 403
				loginResMes.Error = err.Error()
			default:
				loginResMes.Code = 505
				loginResMes.Error = "服务器内部错误"
		}			
	}else{
		loginResMes.Code = 200
		//将登陆成功的用户的ID赋给
		this.UserId = loginMes.UserId
		//登陆成功后将用户加入在线用户列表
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUsers(loginMes.UserId)
		
		//遍历将userMgr中的在线用户加入ResMes中
		for id,_ := range userMgr.onlineUsers{
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
	}
	fmt.Println(user)

	//将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("LoginResMes序列化失败 error = ",err)
		return
	}

	resMes.Data = string(data)
	//将resMes 序列化，准备发送
	data1, err := json.Marshal(resMes)
	if err != nil{
		fmt.Println("ResMes序列化失败 error = ",err)
		return
	}

	//发送数据，将其封装到writePkg函数中
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data1)
	if err != nil {
		fmt.Println("writePkg error = ",err)
	}
	return

}

//处理注册
func (this *UserProcess)ServerProcessRegister(mes *message.Message)(err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil{
		fmt.Println("反序列化失败 err = ",err)
		return 
	}
	//声明一个Message
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//再声明一个LoginResMes
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil{
		if err == model.ERROR_USER_EXIST{
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXIST.Error()
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "注册时发生未知错误"
		}
	}else{
		registerResMes.Code = 200
	}

	//将registerResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil{
		fmt.Println("序列化失败 error = ",err)
		return
	}

	resMes.Data = string(data)
	//将resMes 序列化，准备发送
	data1, err := json.Marshal(resMes)
	if err != nil{
		fmt.Println("ResMes序列化失败 error = ",err)
		return
	}

	//发送数据，将其封装到writePkg函数中
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data1)
	if err != nil {
		fmt.Println("writePkg error = ",err)
	}
	return

}

//通知所有(其他的)在线用户
func (this *UserProcess) NotifyOthersOnlineUsers(userId int){
	for id,up := range userMgr.onlineUsers{
		//过滤掉自己
		if id == userId {
			continue
		}else{
			//通知其他在线用户
			up.NotifyMeOnline(userId)
		}
	}
}

func (this *UserProcess) NotifyMeOnline(userId int){
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline 

	data,err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	mes.Data = string(data)

	data1,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	//发送
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data1)
	if err != nil {
		fmt.Println("notifyOnline失败")
		return
	}
}
