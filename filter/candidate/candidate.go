package candidate

//Candidate 候选人
type Candidate struct {
	Gender     string
	Age        int
	Salary     int
	Experience int
}

//GetGender 读取性别
//Return
//   string 性别 male 为男性，female 为女性
func (c Candidate) GetGender() string {
	return c.Gender
}

//GetAge 读取年龄
//Return
//    int 年龄
func (c Candidate) GetAge() int {
	return c.Age
}

//GetSalary 读取薪资
//Return
//    int 薪资信息
func (c Candidate) GetSalary() int {
	return c.Salary
}

//GetExperience 读取工作经验
//Return
//    int 工作经验
func (c Candidate) GetExperience() int {
	return c.Experience
}
