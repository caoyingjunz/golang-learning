package main

import (
	"fmt"
)

func main() {
	var num float64
	flag := false
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	if num < 0 {
		flag = true
		num = -num
	}

	result := getCubicValue(num)
	if flag {
		fmt.Printf("%.1f\n", -result)
	} else {
		fmt.Printf("%.1f\n", result)
	}
}

func getCubicValue(num float64) float64 {
	min := 0.0
	max := num
	if num < 1 {
		max = 1
	}

	for max-min > 0.00000001 {
		mid := (max + min) / 2
		if mid*mid*mid > num {
			max = mid
		} else if mid*mid*mid < num {
			min = mid
		} else {
			return mid
		}
	}
	return max
}
