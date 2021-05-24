package main

import (
	"bufio"
	"fmt"
	"os"
)

// A~Z 65~90
// a~z 97~122
// ' ' 32

func main() {
	data, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	var tmp = 0
	for index := len(data) - 1; index >= 0; index-- {
		if data[index] < 65 || data[index] > 122 || data[index] < 97 && data[index] > 90 {
			fmt.Printf("%s ", string(data[index+1:index+tmp+1]))
			tmp = 0
		} else {
			tmp++
			if index == 0 {
				fmt.Printf("%s", string(data[:tmp]))
			}
		}
	}

}
