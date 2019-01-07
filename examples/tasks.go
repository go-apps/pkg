package main

import (
	"fmt"
	"time"

	"github.com/go-apps/pkg/core"
)

type HttpServer struct {
	core.ApplicationTask
	notifications chan []string
}

func (s1 *HttpServer) Exec() {
	s1.ApplicationTask.Exec()
	fmt.Println("HttpServer exec.")
	for action := range s1.notifications {
		fmt.Printf("%s", action)
	}
	fmt.Println("HttpServer exec end.")
}

func (s1 *HttpServer) Close() {
	s1.ApplicationTask.Close()
	fmt.Println("HttpServer Close.")
	close(s1.notifications)
	fmt.Println("HttpServer Close end.")
}

type StunServer struct {
	core.ApplicationTask
	notifications chan []string
}

func (s2 *StunServer) Exec() {
	s2.ApplicationTask.Exec()
	fmt.Println("StunServer exec.")

	time.Sleep(5 * time.Second)
	panic("Panic, omg ...")

	for action := range s2.notifications {
		fmt.Printf("%s", action)
	}
	fmt.Println("StunServer exec end.\n")
}

func (s2 *StunServer) Close() {
	s2.ApplicationTask.Close()
	fmt.Println("StunServer Close.")
	close(s2.notifications)
	fmt.Println("StunServer Close end.")
}
