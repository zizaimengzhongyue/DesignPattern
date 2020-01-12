//Description
//每个候选人来面试的时候需要一个编号，这个编号不能重复。所以所有候选人多需要在同一个地方领取编号
package main

import (
	"fmt"

	"./sington"
)

func main() {
	interview01 := sington.GetInstance()
	fmt.Println(interview01.GetID())

	interview02 := sington.GetInstance()
	fmt.Println(interview02.GetID())

	fmt.Printf("%p\n", interview01)
	fmt.Printf("%p\n", interview02)
}
