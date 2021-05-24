package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		var buf bytes.Buffer
		a := strings.TrimSpace(in.Text())
		in.Scan()
		b := strings.TrimSpace(in.Text())

		carry := 0
		for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry != 0; {
			x, y := 0, 0
			if i < 0 {
				x = 0
			} else {
				x = int(a[i] - '0')
				i--
			}
			if j < 0 {
				y = 0
			} else {
				y = int(b[j] - '0')
				j--
			}
			sum := x + y + carry
			buf.WriteByte(byte(sum%10 + '0'))
			carry = sum / 10
		}
		res := []rune(buf.String())
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
		fmt.Println(string(res))
	}
}
