package main

import (
	"fmt"
	"os/exec"
)

func main() {

	fullArgs := make([]string, 0)
	fullArgs = append(fullArgs, []string{"-al"}...)

	output, _ := exec.Command("ls", fullArgs...).CombinedOutput()
	fmt.Println(string(output))

}
