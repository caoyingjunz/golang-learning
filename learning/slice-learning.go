package main

import "fmt"

func main() {
	args := make([]string, 0)
	t1 := append(args, []string{"test1"}...)
	fmt.Println(args[:0])
	args = append(args, []string{"test1"}...)
	t2 := append(args[:0], []string{"test2"}...)
	fmt.Println(args[:0])

	fmt.Println("t1", t1, "t2", t2)
}
