package main

import (
	"fmt"
)

func main() {
	for {
		var input string
		n, err := fmt.Scan(&input)
		if n < 1 || err != nil {
			break
		}

		//fmt.Println(input)
		var tmp int = -1
		slice := []byte(input)
		for i := 0; i < len(slice)-1; i++ {
			for j := i + 1; j < len(slice); j++ {
				if slice[i] == slice[j] {
					slice[j] = 0
					tmp = i
				}
			}

			if tmp != -1 {
				slice[tmp] = 0
			}
		}

		var output byte
		for _, output = range slice {
			if output != 0 {
				fmt.Printf("%c\n", output)
				break
			}
		}

		if output == 0 {
			fmt.Printf("-1\n")
		}

	}
}
