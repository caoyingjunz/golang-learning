package main

import "sync"

// https://blog.csdn.net/m0_37579159/article/details/79306845?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.channel_param

var m sync.Map

func main() {

	go func() {
		m.Store("key", "test1") // write
	}()

	go func() {
		v, ok := m.Load("key") // read
	}()
}
