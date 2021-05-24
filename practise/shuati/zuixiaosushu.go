package main

import "fmt"

func main() {

	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}

		left := n / 2
		right := n / 2

		for left >= 2 {
			if isPrime(left) && isPrime(right) {
				fmt.Println(left)
				fmt.Println(right)
				break
			} else {
				left--
				right++
			}
		}
	}
}

//判断一个数是否为素数
func isPrime(i int) bool {
	count := 0
	for j := 2; j < i; j++ {
		if i%j == 0 {
			count++
		}
	}
	if count == 0 {
		return true
	}
	return false
}
