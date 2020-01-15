package administration

import "fmt"

type HR struct{}

func (h HR) Support() {
	fmt.Println("我们是 HR，我们提供人力相关服务")
}
