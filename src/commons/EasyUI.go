package commons

type Datagrid struct {
	Rows  interface{} `json:"rows"`  //当前页显示的数据
	Total int         `json:"total"` //总个数
}

// EasyUITree tree
type EasyUITree struct {
	Id    int    `json:"id"`
	Text  string `json:"text"`
	State string `json:"state"`
}
