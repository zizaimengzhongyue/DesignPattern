package candidate

import (
	"fmt"
)

type Salesman struct {
}

func (s Salesman) Pass() {
	fmt.Println("您丰富的工作经验给我们留下了深刻的印象，欢迎加入销售部")
}

func (s Salesman) Reject() {
	fmt.Println("您丰富的工作经验给我们留下了深刻的印象，但是跟销售部当前需求不符，希望以后有合作的机会")
}
