package main

import (
	"fmt"
	"time"

	"k8s.io/utils/clock"
)

type TestController struct {
	queue []string
}

func (tc *TestController) AddItems() {
	for {
		time.Sleep(1 * time.Second)
		tc.queue = append(tc.queue, time.Now().String())
	}
}

func (tc *TestController) processNextWorkItem() bool {

	if len(tc.queue) == 0 {
		fmt.Println(tc.queue)
		return false
	}

	//item := tc.queue[0]
	////tc.queue = tc.queue[1:]
	//
	//fmt.Println(item)

	fmt.Println("process")
	return true
}

func (tc *TestController) worker() {
	for tc.processNextWorkItem() {
	}
}

func Until(f func(), stopCh <-chan struct{}) {
	var t clock.Timer
	for {
		select {
		case <-stopCh:
			return
		default:
		}

		func() {
			f()
		}()

		select {
		case <-stopCh:
			return
		case <-t.C():
		}
	}
}

func main() {
	stopCh := make(chan struct{})

	tc := TestController{
		queue: make([]string, 0),
	}

	go tc.AddItems()

	for i := 0; i < 5; i++ {
		go Until(tc.worker, stopCh)
	}

	<-stopCh
}
