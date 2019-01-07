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

type ApplicationTask struct {
	NeedRestart bool
	status      TaskStatus
}

func (t *ApplicationTask) RestartOnCrash() bool {
	return t.NeedRestart
}

func (t *ApplicationTask) Exec() {
	t.status = TaskStatusRunning
}

func (t *ApplicationTask) Close() {
	t.status = TaskStatusClosing
}

func (t *ApplicationTask) Status() TaskStatus {
	return t.status
}
