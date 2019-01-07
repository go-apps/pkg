package core

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-apps/pkg/core/standardpaths"
	"github.com/go-apps/pkg/crashdmp"
)

// coreApplication coreApplication
type coreApplication struct {
	sync.WaitGroup

	inited           bool
	tasks            []IApplicationTask
	ApplicationName  string
	OrganizationName string
	AppDataPath      string
}

var gAPP *coreApplication

// CoreApplication singleton
func CoreApplication() *coreApplication {
	if gAPP == nil {
		gAPP = new(coreApplication)
	}

	return gAPP
}

func init() {
	CoreApplication()
}

func (d *coreApplication) appendOrganizationAndApp() string {
	return fmt.Sprintf("%s/%s", d.OrganizationName, d.ApplicationName)
}

func (d *coreApplication) createAppDataPath() string {
	datapath, err := standardpaths.WritableLocation(standardpaths.AppDataLocation)
	if err != nil {
		datapath = os.TempDir()
	}

	fmt.Printf("datapath=%s\n", datapath)

	d.AppDataPath = fmt.Sprintf("%s/%s", datapath, d.appendOrganizationAndApp())

	if err = os.MkdirAll(d.AppDataPath, os.ModeDir); err != nil {
		fmt.Printf("mkdir error:%+v", err)
	}

	fmt.Printf("appdatapath=%s\n", d.AppDataPath)
	return d.AppDataPath
}

// Initialize 初始化
func (d *coreApplication) Initialize(orgName string, appName string) {
	if d.inited {
		fmt.Println("coreApplication has Initialized.")
		return
	}
	fmt.Println("coreApplication Initialize.")

	d.OrganizationName = orgName
	d.ApplicationName = appName

	var rootDir = d.createAppDataPath()
	dmpPath := fmt.Sprintf("%s/dmps", rootDir)

	os.MkdirAll(dmpPath, os.ModeDir)
	crashdmp.SetupDumpStackTrap(dmpPath)

	d.inited = true
}

func runTask(task IApplicationTask) {
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
func (d *coreApplication) AddTask(task IApplicationTask) {
	fmt.Printf("coreApplication AddTask.\n")
	d.Add(1)
	d.tasks = append(d.tasks, task)
	go func() {
		runTask(task)

		d.Done()
	}()
}

// WaitForShutdown 添加一个应用级协程任务
func (d *coreApplication) WaitForShutdown() {
	fmt.Printf("coreApplication WaitForShutdown.\n")
	for _, task := range d.tasks {
		task.Close()
	}
	d.Wait()
	fmt.Printf("coreApplication WaitForShutdown End.\n")
}

func (d *coreApplication) SetApplicationName(appName string) {
	d.ApplicationName = appName
}

func (d *coreApplication) SetOrganizationName(orgName string) {
	d.OrganizationName = orgName
}

// Run Run
func Run(mainFunc func()) int {
	if !CoreApplication().inited {
		return 1
	}
	exitChan := make(chan int)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		exitChan <- 1
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	mainFunc()

	<-exitChan
	CoreApplication().WaitForShutdown()

	fmt.Printf("Application Quit.\n")

	return 0
}
