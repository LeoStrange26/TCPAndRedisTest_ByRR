package model

import (
	"RedisAndTCP_Project/common/message"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//dao : data access object

type UserDao struct {
	Pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) *UserDao {
	userDao := &UserDao{
		Pool: pool,
	}
	return userDao
}

//在服务器启动时就创建一个UserDao实例（全局）
//在后续操作redis时直接使用即可
var (
	MyUserDao *UserDao
)

//根据用户id返回一个User信息
func (this *UserDao)GetUserById(conn redis.Conn,id int)(user message.User, err error){
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil{
			//在数据库中没有找到相应的id
			err = ERROR_USER_NOTEXIST
		}
		return 
	}
	//将res反序列化
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Printf("反序列化失败 error = %v\n",err)
		return
	}
	return
}

//完成登陆的校验
func (this *UserDao) Login(userId int, userPwd string) (user message.User, err error) {
	//先从中取出一个连接
	conn := this.Pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn,userId)
	if err != nil {
		return 
	}

	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}

	return
}

//完成注册
func (this *UserDao) Register(user *message.User) (err error) {
	//先从中取出一个连接
	conn := this.Pool.Get()
	defer conn.Close()

	_,err = this.GetUserById(conn,user.UserId)
	if err == nil {
		err = ERROR_USER_EXIST
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("序列化出错 error = ",err)
		return
	}
	
	//注册信息入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册信息错误")
		return
	}
	return

}