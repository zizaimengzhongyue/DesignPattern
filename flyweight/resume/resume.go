package resume

import "fmt"

//Resume 简历接口
type Resume interface {
	Description()
}

//ProgrammerResume 程序员简历
type ProgrammerResume struct {
	Name       string
	Age        int
	Experience int
}

//SetName 设置程序员简历姓名
//Params
//    name string 姓名
func (pr *ProgrammerResume) SetName(name string) {
	pr.Name = name
}

//SetAge 设置程序员简历姓名
//Params
//    age int 年龄
func (pr *ProgrammerResume) SetAge(age int) {
	pr.Age = age
}

//SetExperience  设置工作经验
//Params
//    experience int 工作经验
func (pr *ProgrammerResume) SetExperience(experience int) {
	pr.Experience = experience
}

//Description 简历信息
func (pr ProgrammerResume) Description() {
	fmt.Printf("我的姓名是 %s，今年 %d 岁，有 %d 年工作经验\n", pr.Name, pr.Age, pr.Experience)
}
