package model

import (
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"go_code/chatroom1.1/common/message"
)

//我们在服务器启动后，就初始化一个userDao实例
//把它做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao {
		pool : pool,
	}
	return
}


//1. 根据用户id返回一个user实例
func (this *UserDao)getUserById(conn redis.Conn, id int) (user *User, err error) {

	//因为conn.Do()函数返回的值是一个空接口类型的数据，所以要转换成字符串类型
	 res, err := redis.String(conn.Do("hget", "users", id))
	 if err != nil {
		 //要是在users哈希中，没有找到对应的id，就会产生redis.ErrNil的错误
		 if err == redis.ErrNil {
			 err = ERROR_USER_NOTEXITS
		 }
		 return 
	 }
	 user = &User{}
	 //吧res反序列化成User实例
	 err = json.Unmarshal([]byte(res), user)
	 if err != nil {
		 fmt.Println("json.Unmarshal err = ", err)
		 return
	 }
	 return
}

//完成登录校验
//1. Login完成对用户的验证
//2. 如果用户的id和pwd都正确，则返回一个user实例
//3. 如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()//延时关闭

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//进行验证
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//完成注册校验
func (this *UserDao) Register(user *message.User) (err error) {
	//先从UserDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()//延时关闭

	_, err = this.getUserById(conn, user.UserId) 
	if err == nil {
		err = ERROR_USER_EXITS //用户已存在
		return
	}

	//若用户不存在，则可以进行账号数据入库注册操作
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库+
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误err = ", err)
	}
	return
}