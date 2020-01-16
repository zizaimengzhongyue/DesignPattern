//Description
//    通过姓名访问并创建简历，多次访问读取同一份数据。因为数据是同一份，所以对数据的修改会同时生效。
//    类似于在数据库中间增加缓存的一个机制
package main

import (
	"./resume"
)

func main() {
	factory := resume.Factory{MapResume: map[string]*resume.ProgrammerResume{}}

	resume := factory.GetInstance("张三")
	resume.SetAge(25)
	resume.SetExperience(3)

	resume = factory.GetInstance("李四")
	resume.SetAge(30)
	resume.SetExperience(5)

	resume01 := factory.GetInstance("张三")
	resume01.Description()

	resume02 := factory.GetInstance("张三")
	resume02.SetName("王五")
	resume02.SetAge(35)
	resume02.SetExperience(10)

	resume01.Description()
	resume02.Description()
}
