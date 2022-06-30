package desc

// Insert 新增
func Insert(t TbItemDesc) int {
	return insertDescDao(t)
}

// SelByIdService 查询
func SelByIdService(id int) *TbItemDesc {
	return selByIdDao(id)
}
