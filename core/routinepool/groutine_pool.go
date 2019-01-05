package routinepool

type IRunnable interface {
	Run()
	Terminate()
}

type GRoutinePool struct {
	maxThreadCount int
	routines       []IRunnable
}

var globalInstance *GRoutinePool

func GetGlobalInstance() *GRoutinePool {

	return globalInstance
}

func (p *GRoutinePool) SetMaxThreadCount(maxThreadCount int) {

	p.maxThreadCount = maxThreadCount
}

func (p *GRoutinePool) ActiveThreadCount() int {

	return len(p.routines)
}

func (p *GRoutinePool) Start(runnable IRunnable, priority int) {
	go runnable.Run()
}

func (p *GRoutinePool) Cancel(runnable IRunnable) {
	runnable.Terminate()
}
