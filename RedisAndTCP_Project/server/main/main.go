package main

import (
	"RedisAndTCP_Project/server/model"
	"RedisAndTCP_Project/server/processor"
	"fmt"
	"net"
	"time"
)

//处理与客户端的连接
//net.Conn是引用类型（所以不需要加*）
func process(conn net.Conn) {
	defer conn.Close()

	processor := &processor.Processor{
		Conn : conn,
	}
	err := processor.Process()
	if err != nil {
		fmt.Println("服务器与客户端通讯协程错误 error = ",err)
		return
	}
}

//初始化UserDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(model.Pool)
}

func main() {
	model.InitPool("localhost:6379",16,0,300 * time.Second)
	initUserDao()
	// fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp",":8889")
	if err != nil {
		fmt.Println("net listen error = ", err)
		return
	}
	defer listen.Close()

	//一旦连接成功，则始终保持监听状态
	for{
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen connect error = ", err)
			return
		}else{
			//一旦连接成功，开启一个协程进行处理
			go process(conn)
		}
	}
}