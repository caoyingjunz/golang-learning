package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var s string
		n, _ := fmt.Scanln(&s)
		if n == 0 {
			break
		}
		fmt.Println("input:", s)
		for i := 0; i < len(s); i++ {
			if s[i] == '-' {
				if i == 0 || ((s[i-1] < '0' || s[i-1] > '9') && s[i-1] != ')') {
					var j int
					for j = 1; i+j < len(s); j++ {
						if s[i+j] < '0' || s[i+j] > '9' {
							break
						}
					}
					s = s[:i] + "(0-" + s[i+1:i+j] + ")" + s[i+j:]
				}
			}
		}
		//fmt.Print(s)
		arr := make([]string, 0)
		i := 0
		for i < len(s) {
			var temp string
			for i < len(s) && s[i] <= '9' && s[i] >= '0' {
				temp += string(s[i])
				i++
			}
			if temp != "" {
				arr = append(arr, temp)
			}
			i++
		}
		//fmt.Print(arr, len(arr))
		iarr := 0
		op := make([]byte, 100)
		iop := 0
		nums := make([]string, 100)
		inums := 0
		//        flag := false
		for i = 0; i < len(s); i++ {
			//fmt.Print(i, " ")
			if s[i] == '(' || s[i] == '*' || s[i] == '/' {
				op[iop] = s[i]
				iop++
			} else if s[i] == '+' || s[i] == '-' {
				for iop > 0 && op[iop-1] != '(' {
					nums[inums] = string(op[iop-1])
					inums++
					op[iop-1] = s[i]
					iop--
				}
				op[iop] = s[i]
				iop++
				//fmt.Print("+++++++++",op)
			} else if s[i] == ')' {
				for op[iop-1] != '(' {
					nums[inums] = string(op[iop-1])
					inums++
					iop--
				}
				iop--
			} else {
				for i < len(s) && s[i] >= '0' && s[i] <= '9' {
					i++
				}
				nums[inums] = arr[iarr]
				inums++
				iarr++
				i--
			}
		}
		//fmt.Print(iop)
		for iop > 0 {
			nums[inums] = string(op[iop-1])
			//fmt.Print(string(op[iop - 1]))
			iop--
			inums++
		}
		cacl := make([]int, 100)
		icacl := 0
		for i := 0; i < inums; i++ {
			//fmt.Print(nums[i], " ")
			if nums[i] != "+" && nums[i] != "-" && nums[i] != "*" && nums[i] != "/" {
				temp, _ := strconv.Atoi(nums[i])
				cacl[icacl] = temp
				icacl++
			} else {
				res := 0
				if nums[i] == "+" {
					res = cacl[icacl-2] + cacl[icacl-1]
				} else if nums[i] == "-" {
					res = cacl[icacl-2] - cacl[icacl-1]
				} else if nums[i] == "*" {
					res = cacl[icacl-2] * cacl[icacl-1]
				} else {
					res = cacl[icacl-2] / cacl[icacl-1]
				}
				//fmt.Print(res, " ")
				cacl[icacl-2] = res
				icacl--
			}
		}
		fmt.Println(cacl[0])
	}
}
