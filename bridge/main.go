//Description
//    针对一个岗位安排面试之前得先发招聘，针对不同的职位有不同的工作描述
//    创建 JobDescription 这个接口，在 Interview 中使用这个接口的实体类发布工作描述
//    感觉上跟策略模式很像，也可能是我的理解有问题
package main

import (
	"./interview"
	"./jd"
)

func main() {
	interv01 := interview.Interview{Age: 25, Salary: 5000, Experience: 3, JD: jd.Salesman{}}
	interv01.Advertise()

	interv02 := interview.Interview{Age: 25, Salary: 8000, Experience: 3, JD: jd.Programmer{}}
	interv02.Advertise()
}
