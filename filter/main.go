//Descpiton
//    现在有一批候选人名单，我们需要按照年不同的条件对候选人进行过滤，通过创建不同的过滤器，可以筛选出符合条件的
package main

import (
	"fmt"

	"./candidate"
	"./filter"
)

func main() {
	candidates := []candidate.Candidate{
		{Age: 25, Salary: 8000, Experience: 5},
		{Age: 45, Salary: 8000, Experience: 5},
		{Age: 25, Salary: 3000, Experience: 5},
		{Age: 25, Salary: 10000, Experience: 5},
		{Age: 50, Salary: 5000, Experience: 5},
		{Age: 25, Salary: 5000, Experience: 5},
	}

	young := filter.Young{}
	fmt.Println("young:")
	fmt.Println(young.Filter(candidates))

	cheap := filter.Cheap{}
	fmt.Println("cheap:")
	fmt.Println(cheap.Filter(candidates))
}
