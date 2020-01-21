package business

import "fmt"

//IT IT 部门
type IT struct{}

//DoProcessing 提供服务
func (i *IT) DoProcessing() {
	fmt.Println("我们提供 IT 相关服务")
}
