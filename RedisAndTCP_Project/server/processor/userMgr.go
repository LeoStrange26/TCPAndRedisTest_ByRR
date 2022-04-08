package processor

import "fmt"

//UserMgr实例在服务端有且只有一个（userMgr）
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

var (
	userMgr *UserMgr
)

func init() {
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess,1024),
	}
}

//对onlineUser的添加
func (this *UserMgr)AddOnlineUser(up *UserProcess){
	this.onlineUsers[up.UserId] = up
}

//对onlineUser的删除
func (this *UserMgr)DeleteOnlineUser(userId int){
	delete(this.onlineUsers, userId)
}

//查询所有当前在线用户
func (this *UserMgr)GetAllOnlineUsers() map[int]*UserProcess{
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr)GetOnlineUserById(userId int) (up *UserProcess,err error){
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 当前不在线",userId)
		return
	}
	return 
}