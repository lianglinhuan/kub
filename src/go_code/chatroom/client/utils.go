package main
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/common/message"
	_"errors"
	_"io"
)

//定义一个读取客户端数据的函数
func readPkg(conn net.Conn) (mes message.Message, err error) {
	//1.获取客户端发送信息的数据字节长度
	buf := make([]byte, 8096)
	fmt.Println("等待读取客户端的信息。。。。")
	//这里conn.Raed在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn，则就不会阻塞，要是不阻塞，程序就会往下执行
	//所以在客户端关闭前来个延时，等待服务器处理好数据
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		//自定义错误
		// err = errors.New("read pkg header error")
		return
	}
	// fmt.Println("读取到的buf = ", buf[:4])

	//2.根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//3.根据pkgLen读取消息内容
	//这里做一个特殊说明！！！！，conn.read()函数不是从buf[:pkgLen]读取数据的，
	//而是从conn读取数据后，在将读取到的数据信息传入buf[:pkgLen]切片里面的
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err = ", err)
		// err = errors.New("read pkg body error")
		return
	}

	//4.将获取到的数据进行序列化
	//4.1 先将数据转换成message.Message的结构体数据类型
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return

}

