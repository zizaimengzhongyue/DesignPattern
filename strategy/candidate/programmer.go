package candidate

import (
	"fmt"
)

type Programmer struct {
}

func (p Programmer) Pass() {
	fmt.Println("您优秀的专业技能给我们留下了深刻的印象，欢迎加入技术部")
}

func (p Programmer) Reject() {
	fmt.Println("您优秀的专业技能给我们留下了深刻的印象，但是与技术部当前需求不符合，希望以后合作的机会")
}
