package main

import (
	"fmt"
	"time"

	"github.com/go-apps/pkg/core"
)

type HttpServer struct {
	core.Task
	notifications chan []string
}

func (s1 *HttpServer) Exec() {
	s1.Task.Exec()
	fmt.Println("HttpServer exec.")
	for action := range s1.notifications {
		fmt.Printf("%s", action)
	}
	fmt.Println("HttpServer exec end.")
}

func (s1 *HttpServer) Close() {
	s1.Task.Close()
	fmt.Println("HttpServer Close.")
	close(s1.notifications)
	fmt.Println("HttpServer Close end.")
}

type StunServer struct {
	core.Task
	notifications chan []string
}

func (s2 *StunServer) Exec() {
	s2.Task.Exec()
	fmt.Println("StunServer exec.")

	time.Sleep(5 * time.Second)
	panic("Panic, omg ...")

	for action := range s2.notifications {
		fmt.Printf("%s", action)
	}
	fmt.Println("StunServer exec end.\n")
}

func (s2 *StunServer) Close() {
	s2.Task.Close()
	fmt.Println("StunServer Close.")
	close(s2.notifications)
	fmt.Println("StunServer Close end.")
}

// type Apprtc struct {
// 	core.CoreApplication
// }

func Main() {
	var s1 core.IApplicationTask
	s1 = &HttpServer{notifications: make(chan []string)}
	var s2 core.IApplicationTask
	s2 = &StunServer{Task: core.Task{NeedRestart: true}, notifications: make(chan []string)}
	core.CoreApplication().AddTask(s1)
	core.CoreApplication().AddTask(s2)
}

func main() {
	// var app core.IApplication
	// app = &Apprtc{}

	core.CoreApplication().Initialize("XXXX.Ltd", "Test")

	core.Run(Main)
}
