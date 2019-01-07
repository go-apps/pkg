package main

import (
	"github.com/go-apps/pkg/core"
)

func Main() {
	var s1 core.IApplicationTask
	s1 = &HttpServer{notifications: make(chan []string)}
	var s2 core.IApplicationTask
	s2 = &StunServer{ApplicationTask: core.ApplicationTask{NeedRestart: true}, notifications: make(chan []string)}
	core.CoreApplication().AddTask(s1)
	core.CoreApplication().AddTask(s2)
}

func main() {
	core.CoreApplication().Initialize("XXXX.Ltd", "Test")

	core.Run(Main)
}
