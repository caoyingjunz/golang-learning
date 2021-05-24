package main

import "fmt"

// 杨辉三角形变形

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		if n <= 2 {
			fmt.Println(-1)
		} else if n%2 == 1 {
			fmt.Println(2)
		} else if n%4 == 0 {
			fmt.Println(3)
		} else {
			fmt.Println(4)
		}
	}

}
