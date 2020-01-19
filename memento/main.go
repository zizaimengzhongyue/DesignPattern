//Description
//    假设一个写论文的场景，写论文的过程中我们需要频繁备份论文，并在需要的时候可以随时恢复
package main

import (
	"fmt"

	"./memento"
)

func main() {
	careTaker := memento.CareTaker{}

	origination := memento.Originator{}
	origination.SetContent("第一次写入数据")
	careTaker.Add(origination.SaveContentToMemento())
	origination.SetContent("第二次写入数据")
	careTaker.Add(origination.SaveContentToMemento())

	origination.RestoreContentFromMemento(careTaker.Get(1))
	fmt.Println("第二次备份的数据：" + origination.GetContent())
	origination.RestoreContentFromMemento(careTaker.Get(0))
	fmt.Println("第一次备份的数据：" + origination.GetContent())
}
