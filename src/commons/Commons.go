package commons

import (
	"math/rand"
	"strconv"
	"time"
)

// GenId 生成数据库主键
func GenId() int {
	rand.Seed(time.Now().UnixNano())
	id, _ := strconv.Atoi(strconv.Itoa(rand.Intn(10000)) + strconv.Itoa(int(time.Now().Unix())))
	return id
}
