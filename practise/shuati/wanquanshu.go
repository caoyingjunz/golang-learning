package main

import (
	"fmt"
)

func main() {
	for {
		var n int
		fmt.Scanf("%d", &n)
		if n <= 0 {
			break
		}
		var res int
		if n < 6 {
			res = 0
		} else if n >= 6 && n < 28 {
			res = 1
		} else if n >= 28 && n < 496 {
			res = 2
		} else if n >= 496 && n < 8128 {
			res = 3
		} else if n >= 8128 && n < 130816 {
			res = 4
		} else {
			res = 5
		}
		fmt.Println(res)
	}
}
