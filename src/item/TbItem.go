package item

// TbItem 商品
type TbItem struct {
	Id        int
	Title     string
	SellPoint string
	Price     int
	Num       int
	Barcode   string
	Image     string
	Cid       int
	Status    int8
	Created   string
	Updated   string
}

// TbItemChild 给页面使用,实现商品类目
type TbItemChild struct {
	TbItem
	CategoryName string
}

// TbItemDescChild 给修改页面使用
type TbItemDescChild struct {
	TbItem
	CategoryName string
	Desc         string
}
