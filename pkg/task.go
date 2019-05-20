package root

//Task 更改transid传输标志
type Task struct {
	Name string
	ID   string
}

// Record 更新记录
type Record struct {
	KeyMaps map[string]string
	Table   string
	Flags   []string
}

//Tasker 更改Id接口
type Tasker interface {
	UpdateID(task Task) error
	UpdateFlag(record Record) error
}
