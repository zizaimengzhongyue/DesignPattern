//Description
//    创建一个简历列表，里面包含多份简历，可以通过 HasNext 和 Next 方法遍历这份列表
package main

import (
	"fmt"

	"./resume"
)

func getResumes() *resume.Iterator {
	resumes := []*resume.Impl{}
	resume01 := &resume.Impl{Name: "张三"}
	resumes = append(resumes, resume01)
	resume02 := &resume.Impl{Name: "李四"}
	resumes = append(resumes, resume02)
	return &resume.Iterator{List: resumes}
}

func main() {
	resumes := getResumes()
	for resumes.HasNext() {
		fmt.Println(resumes.Next().GetName())
	}
}
