package user
// TbUser 对应数据库中的用户表
type TbUser struct {
	//属性首字母大写:1. 要转换为json   2. 可能出现跨包访问
	Id       int64  //状态：200表示成功
	Username string //返回的数据
	Password string //返回消息
	Phone    string
	Email    string
	Created  string
	Updated  string
}
