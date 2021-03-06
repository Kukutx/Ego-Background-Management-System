package cat

import (
	"commons"
)

// ShowCatByIdService 根据id查询类目
func ShowCatByIdService(id int) *TbItemCat {
	return selByIdDao(id)
}

func showCatByPidService(pid int) (tree []commons.EasyUITree) {
	cats := selByPid(pid)
	tree = make([]commons.EasyUITree, 0)
	for _, n := range cats {
		state := "open"
		if n.IsParent {
			state = "closed"
		}
		tree = append(tree, commons.EasyUITree{n.Id, n.Name, state})
	}
	return
}
