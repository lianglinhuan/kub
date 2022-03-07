package utils

import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom1.1/common/message"
	"io"
)

//在这里将这些方法关联到结构体
type Transfer struct {
	//f分析应该需要哪些字段
	Conn net.Conn
	Buf [8096]byte
}


//定义一个读取服务端数据的函数
func (this *Transfer)ReadPkg() (mes message.Message, err error) {
	//1.获取客户端发送信息的数据字节长度
	//this.Buf := make([]byte, 8096)
	fmt.Println("....等待读取服务端的信息。。。。")
	//这里conn.Raed在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn，则就不会阻塞，要是不阻塞，程序就会往下执行
	//所以在客户端关闭前来个延时，等待服务器处理好数据
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		if err == io.EOF {
			return 
		}else {
			fmt.Println("Conn.Read err = ", err)
			//自定义错误
			// err = errors.New("read pkg header error")
			return 
		}
	}
	// fmt.Println("读取到的buf = ", buf[:4])

	//2.根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//3.根据pkgLen读取消息内容
	//这里做一个特殊说明！！！！，conn.read()函数不是从buf[:pkgLen]读取数据的，
	//而是从conn读取数据后，在将读取到的数据信息传入buf[:pkgLen]切片里面的
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err = ", err)
		// err = errors.New("read pkg body error")
		return 
	}

	//4.将获取到的数据进行序列化
	//4.1 先将数据转换成message.Message的结构体数据类型
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return 
	}
	return

}



func (this *Transfer)WritePkg(data []byte) (err error) {
	//1. 到这时候data就是我们要发送的消息
	//1.1 先把data的长度发送给服务器
	//先获取data的长度，然后将其转换成一个表示长度的byte切片，
	//因为conn.Write([]byte)所传递的数据是[]byte切片
	var pakLen uint32
	pakLen = uint32(len(data))
	// var buf [4]byte
	//binary.BigEndian.PutUint32([]byte, uint32)
	//1.2 该方法可以将uint32类型的数据转换为[]byte类型，并存放在切片当中
	binary.BigEndian.PutUint32(this.Buf[0:4], pakLen)

	//2. 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf[:4]) err = ", err)
		return
	}
	// fmt.Printf("客户端，发送消息的长度为%d\n 内容是%s\n", len(data), string(data))

	//3.发送消息本身
	n, err = this.Conn.Write(data)
	if n != int(pakLen) || err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}
	return
}