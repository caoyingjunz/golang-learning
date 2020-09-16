package main

import (
	"fmt"
)

type Interface interface {
	Hello()
	GoodBye()

}

func Hello(i Interface) {
	i.Hello()
}

func GoodBye(i Interface) {
	i.GoodBye()
}

type Go struct {}

func (g Go) Hello() {
	fmt.Println("Hello, I am Go")
}

func (g Go) GoodBye() {
}

type Python struct {}

func (p Python) Hello() {
	fmt.Println("Hello, I am python")
}

func (p Python) GoodBye() {

}

func main()  {
	g := Go{}
	Hello(g)

    p := Python{}
    Hello(p)

}