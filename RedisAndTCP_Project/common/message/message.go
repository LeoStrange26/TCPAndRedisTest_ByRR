package message

//定义消息类型
const (
	LoginMesType = "LoginMes" 
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType = "SmsMes"
)

//定义用户在线状态的常量
const(
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string  `json:"type"` //消息类型
	Data string  `json:"data"` //消息内容
}

//登陆信息
type LoginMes struct {
	UserId int  `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

//登陆返回信息
type LoginResMes struct {
	Code int `json:"code"`//200 表示登陆成功
	UserIds []int `json:"userIds"` //返回在线用户的id列表
	Error string `json:"error"`//错误信息
}

//注册消息
type RegisterMes struct {
	User User `json:"user"`
}

//注册后的返回信息
type RegisterResMes struct {
	Code int `json:"code"`//200 表示注册成功， 400表示用户已占用
	Error string `json:"error"`//错误信息
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMes struct {
	Content string `json:"content"`
	User
}