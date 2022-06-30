package commons

// EgoResult 客户端服务端数据交互模板
type EgoResult struct {
	Status int         //状态：200表示成功
	Data   interface{} //返回的数据
	Msg    string      //返回消息
}