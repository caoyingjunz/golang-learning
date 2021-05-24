package main

import (
	"fmt"
	"strings"
)

func demo() bool {
	var str string
	var num int
	_, err := fmt.Scan(&str)
	if err != nil {
		return false
	}
	fmt.Scan(&num)
	if num == 0 {
		fmt.Println("")
		return true
	}
	if len(str) <= num {
		fmt.Println(str)
		return true
	}
	if strings.Index(str, "A") >= num || strings.Index(str, "T") >= num {
		fmt.Println(str[:num])
		return true
	}
	if strings.Index(str, "C") == -1 && strings.Index(str, "G") == -1 {
		fmt.Println(str[:num])
		return true
	}
	var cnm, sb, fp int
	for i := 0; i < num; i++ {
		if str[i] == 'C' || str[i] == 'G' {
			cnm++
		}
	}
	fp = cnm
	for i := 1; i < len(str)-num; i++ {
		if str[i-1] == 'C' || str[i-1] == 'G' {
			cnm--
		}
		if str[i+num-1] == 'C' || str[i-1+num] == 'G' {
			cnm++
		}
		if fp < cnm {
			fp = cnm
			sb = i
		}
	}
	fmt.Println(str[sb : sb+num])
	return true
}

func main() {
	for {
		if !demo() {
			return
		}
	}
}
