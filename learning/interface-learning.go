package main

import "fmt"

// https://www.jianshu.com/p/b38b1719636e
// 如果说goroutine和channel是Go并发的两大基石，那么接口是Go语言编程中数据类型的关键。
// 在Go语言的实际编程中，几乎所有的数据结构都围绕接口展开，接口是Go语言中所有数据结构的核心。


type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}

type file struct {
}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("writedata:", data)
	return nil
}

func (f *file) CanWrite() bool {
	return true
}

type Service interface {
	Start()
	Log(string)
}

type Logger struct {
}

func (l *Logger) Log(lg string) {
	fmt.Println(lg)
}

type GameService struct {
	Logger
}

func (g *GameService) Start() {

}

func main() {

	//f := new(file)
	//
	//var writer DataWriter
	//
	//writer = f
	//if writer.CanWrite() {
	//	writer.WriteData("data")
	//}

	var s Service = new(GameService)
	s.Log("ssss")
	s.Start()

}
