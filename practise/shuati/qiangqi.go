package main

import (
	"fmt"
)

func main() {
	var input int
	for {
		if _, err := fmt.Scan(&input); err != nil {
			return
		}

		count := 0
		for i := 7; i <= input; i++ {
			if i%7 == 0 {
				count++
			} else {
				/*
				   for x:=1; i/x > 0; x *= 10 {
				       if i%10 == 7 {
				           count++
				           break
				       }
				   }
				*/
				for n := i; n > 0; n /= 10 {
					if n%10 == 7 {
						count++
						break
					}
				}
			}
		}

		fmt.Println(count)
	}
}
