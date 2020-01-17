package resume

const (
	STATUS_COMMUCATION = 1
	STATUS_INTERVIEW   = 2
	STATUS_PASS        = 3
)

//Resume 简历
type Resume interface {
	GetName() string
	GetStatus() int
}
