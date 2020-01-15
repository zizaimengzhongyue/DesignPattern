package administration

import "fmt"

type IT struct{}

func (i IT) Support() {
	fmt.Println("我们是 IT 部门，我们负责提供 IT 相关的服务")
}
