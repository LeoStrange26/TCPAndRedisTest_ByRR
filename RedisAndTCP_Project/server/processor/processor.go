package processor

//总控
import (
	"RedisAndTCP_Project/common/message"
	"RedisAndTCP_Project/common/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//协程中处理各类请求（根据消息类型）
func (this *Processor)ServerProcessMes( mes *message.Message) (err error){
	fmt.Println("mes=", mes)
	switch mes.Type {
		case message.LoginMesType:
			up := &UserProcess{
				Conn: this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			up := &UserProcess{
				Conn: this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType:
			smsProcess := &SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default:
			fmt.Println("类型不存在")
	}
	return
}

func (this *Processor)Process()(err error) {
	for{
		// fmt.Println("读取客户端发送的数据...")
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		//先读到消息
		mes, err := tf.ReadPkg()
		if err != nil{
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
			}else{
				fmt.Println("读取数据出错")
			}
			return err
		}
		fmt.Println(mes)

		err = this.ServerProcessMes(&mes)
		if err != nil{
			return err
		}
	}
	
}

