package core

type TaskStatus int

const (
	TaskStatusRunning = iota
	TaskStatusClosing
	TaskStatusClosed
)

// IApplicationTask 协程任务
type IApplicationTask interface {
	Exec()
	Close()
	RestartOnCrash() bool
	Status() TaskStatus
}

type Task struct {
	NeedRestart bool
	status      TaskStatus
}

func (t *Task) RestartOnCrash() bool {
	return t.NeedRestart
}

func (t *Task) Exec() {
	t.status = TaskStatusRunning
}

func (t *Task) Close() {
	t.status = TaskStatusClosing
}

func (t *Task) Status() TaskStatus {
	return t.status
}
