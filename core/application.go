package core

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type TaskStatus int

const (
	TaskStatusRunning = iota
	TaskStatusClosing
	TaskStatusClosed
)

// ITask 协程任务
type ITask interface {
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

// IApplication IApplication
type IApplication interface {
	// Initialize 初始化, 只会进程调用一次
	Initialize()

	// AddTask 添加一个应用级协程任务
	AddTask(task ITask)

	// Main 初始化
	Main()

	// 运行
	WaitForShutdown()
}

// CoreApplication CoreApplication
type CoreApplication struct {
	sync.WaitGroup

	tasks            []ITask
	ApplicationName  string
	OrganizationName string
}

// Initialize 初始化
func (d *CoreApplication) Initialize() {
	fmt.Printf("CoreApplication Initialize.\n")
	var rootDir = "d:/"
	d.setupDumpStackTrap(rootDir)
}

func runTask(task ITask) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("task panic：", err)

			if task.RestartOnCrash() && task.Status() == TaskStatusRunning {
				fmt.Println("task restart")
				runTask(task)
			}
		}
	}()
	task.Exec()
}

// AddTask 添加一个应用级协程任务
func (d *CoreApplication) AddTask(task ITask) {
	fmt.Printf("CoreApplication AddTask.\n")
	d.Add(1)
	d.tasks = append(d.tasks, task)
	go func() {
		runTask(task)

		d.Done()
	}()
}

// WaitForShutdown 添加一个应用级协程任务
func (d *CoreApplication) WaitForShutdown() {
	fmt.Printf("CoreApplication WaitForShutdown.\n")
	for _, task := range d.tasks {
		task.Close()
	}
	d.Wait()
	fmt.Printf("CoreApplication WaitForShutdown End.\n")
}

func (d *CoreApplication) SetApplicationName(appName string) {
	d.ApplicationName = appName
}

func (d *CoreApplication) SetOrganizationName(orgName string) {
	d.OrganizationName = orgName
}

func Run(d IApplication) {
	d.Initialize()

	exitChan := make(chan int)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		exitChan <- 1
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	d.Main()

	<-exitChan
	d.WaitForShutdown()

	fmt.Printf("Application Quit.\n")
}
