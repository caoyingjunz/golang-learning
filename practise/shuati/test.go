package main

import (
	"bufio"
	"fmt"
	//	"log"
	"os"
	"strings"
)

func main() {
	//fmt.Println(strings.Fields("hello widuu golang")) //out  [hello widuu golang]

	reader := bufio.NewReader(os.Stdin)

	data, _, _ := reader.ReadLine()
	command := string(data)

	splitArr := strings.Fields(command)
	strLen := len(splitArr)
	if len(splitArr) > 0 {
		fmt.Println(len(splitArr[strLen-1]))

		//log.Println("command", command)
	} else {
		fmt.Println(0)
	}

}
