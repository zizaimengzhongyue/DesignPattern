package interview

//JobDescription 工作描述接口
type JobDescription interface {
	JobDescription(age int, salary int, experience int)
}

//Interview 面试
type Interview struct {
	Age        int
	Salary     int
	Experience int
	JD         JobDescription
}

//Advertise 发布招聘
func (i Interview) Advertise() {
	i.JD.JobDescription(i.Age, i.Salary, i.Experience)
}
