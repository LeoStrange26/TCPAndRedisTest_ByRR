package main

import (
	"fmt"
	"RedisAndTCP_Project/client/processor"
)

func main() {
	var key int   //用户输入
	var loop bool = true //是否继续循环
	for loop{
		showMsg()
		fmt.Scanf("%d\n", &key)
		handleWithKey(key,&loop)
	}
}

var userId int
var userPwd string
var userName string

func handleWithKey(key int, loop *bool) {
	switch key {
		case 1://登陆
			fmt.Println("登陆聊天室")
		
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n",&userPwd)
			up := &processor.UserProcessor{}
			err := up.Login(userId,userPwd)
			if err != nil{
				fmt.Println("登陆出错")
				return
			}
			
		case 2://注册
			fmt.Println("新用户注册")
			
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n",&userPwd)
			fmt.Println("请输入用户昵称")
			fmt.Scanf("%s\n",&userName)
			up := &processor.UserProcessor{}
			up.Register(userId, userPwd,userName)
			
		case 3://退出
			fmt.Println("你选择了退出系统")
			*loop = false
			
		default://其他非法输入
			fmt.Println("你的输入有误,请重新输入")

	}
}

func showMsg(){
	fmt.Println("\t----------欢迎登录多人聊天系统----------")
	fmt.Println("\t\t  1、登陆聊天室")
	fmt.Println("\t\t  2、新用户注册")
	fmt.Println("\t\t  3、退出系统")
	fmt.Println("\t\t   请选择(1-3)")
	fmt.Println("\t--------------------------------------")
}