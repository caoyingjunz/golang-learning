package main

import (
	"fmt"
	//"bufio"
	//"strings"
	//"os"
	"strconv"
)

func main() {
	for {
		var a, b int
		if n, _ := fmt.Scan(&a, &b); n == 0 {
			break
		}

		fmt.Println(getResult(a, b))
	}
}

/*
func Getsum(a,b int) int  {
    if a<0||b<0{
        return 0
    }
    if a==0||a==1||b==1{
        return 1

    }
    return Getsum(a,b-1)+Getsum(a-b,b)
}
*/

/*
func main() {
    for {
        /*
        inputReader := bufio.NewReader(os.Stdin)
        input, _ := inputReader.ReadString('\n') //func (b *Reader) ReadString(delim byte) (line string, err error) ,‘S’ 这个例子里使用S表示结束符，也可以用其它，如'\n'
        input = strings.Trim(input, " ")
        s := strings.Split(input, " ")
        //fmt.Println(s)
        m , _ := strconv.Atoi(s[0])
        //fmt.Println(s[1])
        n, _ := strconv.Atoi(string([]byte{[]byte(s[1])[0]}))
        fmt.Println(getResult(m, n))
*/
/*
        for{
            var a,b int
            if n,_:=fmt.Scan(&a,&b);n==0{
                break
            }
            fmt.Println(getResult(a, b))
        }
    }
}
*/

func getResult(m, n int) int {
	result := 0

	var backtrack func(start, m, n int, ll string)

	backtrack = func(start, m, n int, ll string) {
		if m < 0 {
			return
		}
		if n == 0 && m > 0 {
			return
		}

		if n == 0 || m == 0 {
			result++
			return
		}

		for i := start; i > 0; i-- {
			backtrack(i, m-i, n-1, ll+","+strconv.Itoa(i))
		}
	}

	backtrack(m, m, n, "")
	return result
}
