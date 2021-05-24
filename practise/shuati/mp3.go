/*
MP3光标移动
每页显示4条
*/
package main

import "fmt"

func processMp3(count int, op string) {
	pos := 1
	if count <= 4 {
		for _, v := range op {
			if v == 'U' {
				if pos == 1 {
					pos = count
				} else {
					pos--
				}
			} else if v == 'D' {
				if pos == count {
					pos = 1
				} else {
					pos++
				}
			}
		}

		for i := 1; i <= count; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
		fmt.Println(pos)
		return
	}

	//大于4首
	start := 1
	end := start + 3
	for _, v := range op {
		if v == 'U' {
			if pos == 1 {
				pos = count
				start = count - 3
				end = count
			} else {
				if pos == start {
					start--
					end--
				}
				pos--
			}
		} else {
			if pos == count {
				pos = 1
				start = 1
				end = 4
			} else {
				if pos == end {
					start++
					end++
				}
				pos++
			}
		}
	}
	for i := start; i <= end; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println(pos)
}

func main() {
	count := 0
	op := ""

	for {
		n, err := fmt.Scan(&count)
		if n == 0 || err != nil {
			break
		}
		n, err = fmt.Scan(&op)
		if n == 0 || err != nil {
			break
		}

		processMp3(count, op)
	}
}
