package business

import "fmt"

//HR hr
type HR struct{}

//DoProcessing 提供服务
func (h *HR) DoProcessing() {
	fmt.Println("我们负责提供人力薪资相关服务")
}
