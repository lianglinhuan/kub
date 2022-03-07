package process
import (
	"fmt"
	"go_code/chatroom1.1/common/message"
	"encoding/json"
)

func outputGroupMes(mes *message.Message) {
	//1.反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	//2.显示
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
	fmt.Println()
}