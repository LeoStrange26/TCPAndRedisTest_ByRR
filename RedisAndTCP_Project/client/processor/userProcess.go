package processor

import (
	// "RedisAndTCP_Project/client/model"
	"RedisAndTCP_Project/common/message"
	// "encoding/binary"
	"RedisAndTCP_Project/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	//字段...
}

func (this *UserProcessor)Login(userId int, userPwd string) (err error) {
	//1、连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		//连接出错则直接返回
		fmt.Println("net.Dial error = ", err)
		return
	}
	defer conn.Close()

	//2、准备通过conn向服务器发送消息
	//创建一个Message结构体
	var mes message.Message
	mes.Type = message.LoginMesType

	//3、创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4、将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes Marshal error:", err)
		return
	}
	mes.Data = string(data)

	//5、将mes序列化
	data1, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("Message Marshal error:", err)
		return
	}

	//6、将data发送给服务器
	tf := &utils.Transfer{
		Conn: conn,
		
	}
	tf.WritePkg(data1)


	//7、处理服务器返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg error: ", err)
		return
	}

	//8、将mes的Data反序列化为LoginResMes
	var loginResMes message.LoginResMes
	json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		//登陆成功后初始化Curuser对象
		Curuser.Conn = conn
		Curuser.UserId = userId
		Curuser.UserStatus = message.UserOnline

		//登陆成功后开启一个协程，不断监听服务器发来的消息
		go serverProcessMes(conn)
		fmt.Println("登陆成功")

		//登陆成功显示在线用户id列表
		//并从服务器返回的登陆信息中读取返回的所有在线用户的id，循环将每个在线用户加入客户端维护的onlineUsers中
		for _,id := range loginResMes.UserIds{
			fmt.Println("当前在线用户列表：")
			fmt.Println("用户id:\t",id)
			//完成客户端的onlineUsers初始化
			user := &message.User{
				UserId : id,
				UserStatus: message.UserOnline,
			}
			onlineUsers[id] = user
		}

		fmt.Println()
		for{
			ShowMenu()
		}

	}else{
		fmt.Println(loginResMes.Error)
	}

	return
}

func (this *UserProcessor)Register(userId int, userPwd,userName string) (err error) {
	//1、连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		//连接出错则直接返回
		fmt.Println("net.Dial error = ", err)
		return
	}
	defer conn.Close()

	//2、准备通过conn向服务器发送消息
	//创建一个Message结构体
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3、创建一个RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4、将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("registerMes Marshal error:", err)
		return
	}
	mes.Data = string(data)

	//5、将mes序列化
	data1, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("Message Marshal error:", err)
		return
	}

	//6、将data发送给服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data1)


	//7、处理服务器返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg error: ", err)
		return
	}

	//8、将mes的Data反序列化为RegisterResMes
	var registerResMes message.RegisterResMes
	json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code == 200{
		fmt.Println("注册成功")
		//注册成功后又自动跳回主界面
	}else{
		fmt.Println(registerResMes.Error)
	}

	return

}