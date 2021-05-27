
1. 字符串最后一个单词的长度

描述
计算字符串最后一个单词的长度，单词以空格隔开，字符串长度小于5000。

输入描述：
输入一行，代表要计算的字符串，非空，长度小于5000。

输出描述：
输出一个整数，表示输入字符串最后一个单词的长度。

示例1
输入：
hello nowcoder
复制
输出：
8
复制
说明：
最后一个单词为nowcoder，长度为8


package main

import (
	"fmt"
)

func main() {
	var str, tmp string
	for {
		n, _ := fmt.Scan(&str)
		if n != 0 {
			tmp = str
		}
		if n == 0 {
			fmt.Println(len(tmp))
			break
		}
	}
}

2. 计算某字母出现次数

描述
写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字母，然后输出输入字符串中该字母的出现次数。不区分大小写，字符串长度小于500。

输入描述：
第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字母。

输出描述：
输出输入字符串中含有该字符的个数。

示例1
输入：
ABCabc
A
复制
输出：
2

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var arr []string
    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF {
            break
        } else {
            arr = append(arr, strings.ToLower(string(line)))
        }
    }
    if len(arr[0]) <= 0 {
        fmt.Printf("%v", 0)
        return
    }
    var count int
    for i := 0; i < len(arr[0]); i ++ {
        if arr[0][i] ^ arr[1][0] == 0 {
            count ++
        }
    }
    fmt.Printf("%v", count)
}

3. 明明的随机数

描述
明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，他先用计算机生成了N个1到1000之间的随机整数（N≤1000），对于其中重复的数字，只保留一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。请你协助明明完成“去重”与“排序”的工作(同一个测试用例里可能会有多组数据(用于不同的调查)，希望大家能正确处理)。


注：测试用例保证输入参数的正确性，答题者无需验证。测试用例不止一组。

当没有新的输入时，说明输入结束。

输入描述：
注意：输入可能有多组数据(用于不同的调查)。每组数据都包括多行，第一行先输入随机整数的个数N，接下来的N行再输入相应个数的整数。具体格式请看下面的"示例"。

输出描述：
返回多行，处理后的结果

package main

import(
    "fmt"
    "os"
    "strconv"
    "bufio"
)


//去重和排序
func main(){

    scanner:=bufio.NewScanner(os.Stdin)

    for scanner.Scan(){
        in:=scanner.Text()

        if len(in)==0{
            break
        }
        mark:=make([]bool,1000)

        count,_:=strconv.Atoi(string(in))
        for i:=0;i<count;i++{
            scanner.Scan()
            num,_:=strconv.Atoi(string(scanner.Text()))

            mark[num]=true
        }

        for k,v:=range mark{
            if v==true{
                fmt.Println(k)
            }
        }

    }

4. 字符串分隔

描述
•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。

输入描述：
连续输入字符串(输入多次,每个字符串长度小于100)

输出描述：
输出到长度为8的新字符串数组

示例1
输入：
abc
123456789
复制
输出：
abc00000
12345678
90000000
复制

package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        str := scanner.Text()
        cnt := 0
        var builder strings.Builder

        for i:=0;i<len(str);i++ {
            cnt++
            builder.WriteByte(str[i])
            if cnt == 8 {
                cnt = 0
                fmt.Println(builder.String())
                builder = strings.Builder{}
            }
        }
        if cnt > 0 {
            for i:=cnt;i<8;i++ {
                builder.WriteByte('0')
            }
            fmt.Println(builder.String())
        }

    }
}

5. 进制转换

描述
写出一个程序，接受一个十六进制的数，输出该数值的十进制表示。

输入描述：
输入一个十六进制的数值字符串。注意：一个用例会同时有多组输入数据，请参考帖子https://www.nowcoder.com/discuss/276处理多组输入的问题。

输出描述：
输出该数值的十进制字符串。不同组的测试用例用\n隔开。

示例1
输入：
0xA
0xAA
复制
输出：
10
170
复制

package main

import (
    "fmt"
    "strconv"
)

func main() {
    for {
        var str string
        _, _ = fmt.Scan(&str)
        ret, err := strconv.ParseInt(str, 0, 32) // 字符转数值
        if err != nil {
            break
        }
        fmt.Println(ret)
    }
}


6. 质数因子
描述
功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）

最后一个数后面也要有空格

输入描述：
输入一个long型整数

输出描述：
按照从小到大的顺序输出它的所有质数的因子，以空格隔开。最后一个数后面也要有空格。

示例1
输入：
180
复制
输出：
2 2 3 3 5

package main

import (
    "fmt"
)

func main() {
    var input int
    fmt.Scan(&input)

    for i := 2; i * i <= input; {
        if input % i == 0 {
            fmt.Printf("%d ", i)
            input = input / i
        } else {
            i++
        }
    }
    if input > 1 {
        fmt.Printf("%d ", input)
    }

    fmt.Println()
}

6. 取近似值

描述
写出一个程序，接受一个正浮点数值，输出该数值的近似整数值。如果小数点后数值大于等于5,向上取整；小于5，则向下取整。

输入描述：
输入一个正浮点数值

输出描述：
输出该数值的近似整数值

示例1
输入：
5.5
复制
输出：
6

package main

import (
	"fmt"
)

func main() {
	a := 0.0
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", int(a+0.5))
		}
	}
}

7. 合并表记录

描述
数据表记录包含表索引和数值（int范围的正整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照key值升序进行输出。

输入描述：
先输入键值对的个数
然后输入成对的index和value值，以空格隔开

输出描述：
输出合并后的键值对（多行）

4
0 1
0 2
1 2
3 4

0 3
1 2
3 4


package main

import(
    "bufio"
    "os"
    "strconv"
    "fmt"
    "strings"
)

func main(){
    scan := bufio.NewScanner(os.Stdin)
    for scan.Scan(){
        num,_:= strconv.Atoi(scan.Text())
        var a [1001]int
        for i:=0;i<num;i++{
            scan.Scan()
            nums := strings.Split(scan.Text()," ")
            b,_:= strconv.Atoi(nums[0])
            c,_:= strconv.Atoi(nums[1])
            a[b] += c
        }
        for k,v := range a{
            if v != 0{
                fmt.Printf("%d %d\n",k,v)
            }
        }
    }
}

9. 提取不重复的整数

描述
输入一个int型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
保证输入的整数最后一位不是0。
输入描述：
输入一个int型整数

输出描述：
按照从右向左的阅读顺序，返回一个不含重复数字的新的整数

输入：
9876673
复制
输出：
37689
复制

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var numm = map[int]bool{}
    read := bufio.NewReader(os.Stdin)
    line, _ := read.ReadString('\n')
    line = strings.Trim(line, "\r\n")
    nums := []byte(line)
    for i := len(line) - 1; i >= 0; i-- {
        temp, _ := strconv.Atoi(string(nums[i]))
        if !numm[temp] {
            numm[temp] = true
            fmt.Print(temp)
        }
    }
}

10. 字符个数统计

描述
编写一个函数，计算字符串中含有的不同字符的个数。字符在ACSII码范围内(0~127)，换行表示结束符，不算在字符里。不在范围内的不作统计。多个相同的字符只计算一次
例如，对于字符串abaca而言，有a、b、c三种不同的字符，因此输出3。
输入描述：
输入一行没有空格的字符串。

输出描述：
输出范围在(0~127)字符的个数。

示例1
输入：
abc
复制
输出：
3

package main
import "fmt"

func main(){
    var line string
    fmt.Sprintf("请输入一行字符串:")
    fmt.Scanln(&line)
    maphash := make(map[rune]bool)
    for _,v:=range line{
        if _,ok := maphash[v];!ok{
            maphash[v]=true
        }
    }

    fmt.Println(len(maphash))
}

11. 数字颠倒

描述
输入一个整数，将这个整数以字符串的形式逆序输出
程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001


输入描述：
输入一个int整数

输出描述：
将这个整数以字符串的形式逆序输出

示例1
输入：
1516000
复制
输出：
0006151

package main

import (
	"fmt"
    "strconv"
)

func main() {

	GetReverseString()

}

func GetReverseString() {
	var num int
	fmt.Scanf("%d",&num)
	s:=strconv.Itoa(num)
	end:=len(s)-1
	for i:=end;i>=0	 ;i--  {
		fmt.Printf("%c",s[i])
	}
}

12. 字符串反转
描述
接受一个只包含小写字母的字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

输入描述：
输入一行，为一个只包含小写字母的字符串。

输出描述：
输出该字符串反转后的字符串。

示例1
输入：
abcd
复制
输出：
dcba

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	runes := []rune(text)
	for i := len(runes) - 1; i >= 0; i -- {
        fmt.Print(string(runes[i]))
	}
	fmt.Println("")
}

13. 句子逆序
描述
将一个英文语句以单词为单位逆序排放。例如“I am a boy”，逆序排放后为“boy a am I”
所有单词之间用一个空格隔开，语句中除了英文字母外，不再包含其他字符

输入描述：
输入一个英文语句，每个单词用空格隔开。保证输入只包含空格和字母。

输出描述：
得到逆序的句子

示例1
输入：
I am a boy
复制
输出：
boy a am I

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Input() ([]byte, error) {
	res := []byte{}
	buf := make([]byte, 128)
	var size int
	var err error
	reader := bufio.NewReader(os.Stdin)
	for {
		size, err = reader.Read(buf)
		res = append(res, buf[0:size]...)
		if err != nil {
			break
		}
	}

	return res, err
}

func main() {
	strs, _ := Input()
	lines := bytes.Split(strs, []byte{0xa})
	if len(strs) == 0 || len(lines) < 1 {
		return
	}

	strs1 := bytes.Split(lines[0], []byte{byte(' ')})

	for i := len(strs1) - 1; i > 0; i-- {
		fmt.Printf("%s ", string(strs1[i]))
	}
	fmt.Printf("%s\n", strs1[0])
}

14. 字符串排序

描述
给定n个字符串，请对n个字符串按照字典序排列。
输入描述：
输入第一行为一个正整数n(1≤n≤1000),下面n行为n个字符串(字符串长度≤100),字符串中只含有大小写字母。
输出描述：
数据输出n行，输出结果为按照字典序排列的字符串。
示例1
输入：
9
cap
to
cat
card
two
too
up
boat
boot
复制
输出：
boat
boot
cap
card
cat
to
too
two
up

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sort"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        input := scanner.Text()
        if len(input) == 0{
            break
        }
        n, _ := strconv.Atoi(input)
        strSlice := []string{}
        for i:=0; i<n; i++{
            scanner.Scan()
            input := scanner.Text()
            strSlice = append(strSlice, input)
        }
        sort.Strings(strSlice)

        for _, str := range strSlice{
            fmt.Printf("%s\n", str)
        }
    }
}

15. 求int型正整数在内存中存储时1的个数

描述
输入一个int型的正整数，计算出该int型数据在内存中存储时1的个数。

输入描述：
 输入一个整数（int类型）

输出描述：
 这个数转换成2进制后，输出1的个数

示例1
输入：
5
复制
输出：
2

package main

import (
    "fmt"
)

func main() {
    n := 0
    fmt.Scan(&n)
    counter := 0

    for {
        if n == 0 {
            break
        }
        if n % 2 == 1 {
            counter++
        }
        n = n >> 1

    }

    fmt.Println(counter)
}

16. 购物单

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 核心问题：第i个物品放或者不放
	// 总钱数（价格为10的整数倍，所以可以除10减少循环次数，最后再乘10即可,<32000），物品个数<60
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	nmStr := string(line)
	nmFields := strings.Fields(nmStr)
	n, _ = strconv.Atoi(nmFields[0])
	m, _ = strconv.Atoi(nmFields[1])
	// 对金额作除10处理
	n /= 10
	// 物品价格数组,由于已指定附件个数<=2，则可用二维数组表示物品价格
	weight := [60][3]int{}
	// 物品价值数组，即重要度与价格的乘积为物品价值
	value := [60][3]int{}
	// 表示i个物品在第j钱数时能获得的最大价值
	f := [3200]int{}
	// 先获取weigth和value
	for i := 1; i <= m; i++ {
		var v, p, q int
		line, _, _ = reader.ReadLine()
		vpqStr := string(line)
		vpqFields := strings.Fields(vpqStr)
		v, _ = strconv.Atoi(vpqFields[0])
		p, _ = strconv.Atoi(vpqFields[1])
		q, _ = strconv.Atoi(vpqFields[2])
		// 物品价格也作除10处理
		v /= 10
		if q == 0 {
			// 主件
			weight[i][0] = v
			value[i][0] = v * p
		} else {
			// 判断是否为第一个附件
			if weight[q][1] == 0 {
				weight[q][1] = v
				value[q][1] = v * p
			} else {
				weight[q][2] = v
				value[q][2] = v * p
			}
		}
	}
	// 遍历获取最大价值
	for i := 1; i <= m; i++ {
		for j := n; j >= weight[i][0]; j-- {
			// 筛选掉非主件
			if weight[i][0] <= 0 {
				continue
			}
			// 默认可以容下第i个主件，取放第i个主件和不放第i个主件中总价值较大者
			f[j] = getMax(f[j],f[j-weight[i][0]]+value[i][0])
			// 可以容下第i个主件和第i个主件的第一个附件
			if weight[i][1] > 0 && j >= weight[i][0] + weight[i][1] {
				f[j] = getMax(f[j],f[j-weight[i][0]-weight[i][1]]+value[i][0]+value[i][1])
			}
			// 可以容下第i个主件和第i个主件的第二个附件
			if weight[i][2] > 0 && j >= weight[i][0] + weight[i][2] {
				f[j] = getMax(f[j],f[j-weight[i][0]-weight[i][2]]+value[i][0]+value[i][2])
			}
			// 可以容下第i个主件和第i个主件的全部附件
			if weight[i][1] > 0 && weight[i][2] > 0 && j >= weight[i][0] + weight[i][1] + weight[i][2] {
				f[j] = getMax(f[j],f[j-weight[i][0]-weight[i][1]-weight[i][2]]+value[i][0]+value[i][1]+value[i][2])
			}
		}
	}
	fmt.Println(f[n]*10)
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

17. 坐标移动

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	str, _, err := in.ReadLine()
	if err != nil {
		return
	}

	list := strings.Split(string(str), ";")
	x, y := 0, 0
	for _, v := range list {

		value := []byte(v)

		// 判断坐标有效性
		if len(value) <= 1 {
			continue
		}
		n, err := strconv.ParseInt(string(value[1:]), 10, 64)
		if err != nil {
			continue
		}
		switch string(value[0]) {
		case "A":
			x -= int(n)
		case "D":
			x += int(n)
		case "W":
			y += int(n)
		case "S":
			y -= int(n)
		default:
			continue
		}
	}
	fmt.Printf("%v,%v", x, y)
}

识别有效的IP地址和掩码并进行分类统计

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type interval struct {
	minIP uint
	maxIP uint
	count int
}

func parseIP(ip string) (uint, bool) {
	s := strings.Split(ip, ".")
	if len(s) != 4 {
		return 0, false
	}
	ans := uint(0)
	for i := 0; i < 4; i++ {
		v, err := strconv.Atoi(s[i])
		if err != nil {
			return 0, false
		}
		if v < 0 || v > 255 {
			return 0, false
		}
		ans = ans * uint(256) + uint(v)
	}
	return ans, true
}


func forseParseIP(ip string) uint {
	s := strings.Split(ip, ".")
	ans := uint(0)
	for i := 0; i < 4; i++ {
		v, _ := strconv.Atoi(s[i])
		ans = ans * uint(256) + uint(v)
	}
	return ans
}

func checkMask(umask uint) bool {
	return umask != 0 && umask != 0xFFFFFFFF && (((((^umask) + 1) & umask) - 1) | umask == 0xFFFFFFFF)
}

func find(intervals []*interval, uip uint) bool {
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	intervals := []*interval{
		&interval{
			minIP: forseParseIP("1.0.0.0"),
			maxIP: forseParseIP("126.255.255.255"),
		},
		&interval{
			minIP: forseParseIP("128.0.0.0"),
			maxIP: forseParseIP("191.255.255.255"),
		},
		&interval{
			minIP: forseParseIP("192.0.0.0"),
			maxIP: forseParseIP("223.255.255.255"),
		},
		&interval{
			minIP: forseParseIP("224.0.0.0"),
			maxIP: forseParseIP("239.255.255.255"),
		},
		&interval{
			minIP: forseParseIP("240.0.0.0"),
			maxIP: forseParseIP("255.255.255.255"),
		},
		&interval{
			minIP: forseParseIP("10.0.0.0"),
			maxIP: forseParseIP("10.255.255.255"),
		},
		&interval{
			minIP: forseParseIP("172.16.0.0"),
			maxIP: forseParseIP("172.31.255.255"),
		},
		&interval{
			minIP: forseParseIP("192.168.0.0"),
			maxIP: forseParseIP("192.168.255.255"),
		},
	}
	var invalidCount int

	for {
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		ss := strings.Split(string(s), "~")
		if len(ss) != 2 {
			invalidCount++
			continue
		}
		ip := ss[0]
		mask := ss[1]
		uip, ok := parseIP(ip)
		if !ok {
			invalidCount++
			continue
		}
		umask, ok := parseIP(mask)
		if !ok || !checkMask(umask) {
			invalidCount++
			continue
		}
		for _, itv := range intervals {
			if uip <= itv.maxIP && uip >= itv.minIP {
				itv.count++
			}
		}
	}
	fmt.Printf("%d %d %d %d %d %d %d", intervals[0].count, intervals[1].count, intervals[2].count, intervals[3].count, intervals[4].count,
		invalidCount, intervals[5].count + intervals[6].count + intervals[7].count)
}

简单错误记录
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// var num [2]int
	input := bufio.NewScanner(os.Stdin)
	var output [][2]string
	flag := make(map[[2]string]int)
	for input.Scan() {

		name := ""
		//用数组加map
		line := input.Text()
		key := strings.Split(line, " ")
		for i, v2 := range key {

			if i == 0 {
				//对名字进行处理

				for i := 0; i < len(v2); i++ {

					if v2[i:i+1] == "\\" {
						name = v2[i+1:]
					}
				}

				if len(name) > 16 {
					name = name[len(name)-16:]
				}

			}
			if i == 1 {
				var temp [2]string
				temp[0] = name
				temp[1] = v2
				//行数
				_, yes := flag[temp]
				if yes {

					flag[temp]++

				} else {
					// fmt.Println(v2)

					// fmt.Println(temp)
					output = append(output, temp)
					flag[temp] = 1

				}
			}

		}

	}

	if len(output) > 8 {

		for i := len(output) - 8; i < len(output); i++ {

			fmt.Printf("%s %s %d",output[i][0],  output[i][1],flag[output[i]])
            fmt.Println()
		}
	} else {
		for _, v := range output {

			fmt.Printf("%s %s %d",v[0],  v[1],flag[v])
                        fmt.Println()
		}
	}
}

密码验证合格程序
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	ps := []string{}

	for {
		buf,_,_:= r.ReadLine()
		if len(buf) == 0 {
			break
		}
		ps = append(ps,string(buf))
	}
	for _,passwd := range ps{
		if check(passwd) {
			fmt.Println("OK")
		}else {
			fmt.Println("NG")
		}
	}

}
func check(str string) bool  {
	const (
		nums = "1234567890"
		letters = "qwertyuiopasdfghjklzxcvbnm"
		upLetters = "QWERTYUIOPASDFGHJKLZXCVBNM"
		others = "~`!#$%^&*()-_+={[}]|\\:;,.<>?/\""
	)
	if len(str) <= 8 {
		return false
	}
	count := 0
	if strings.ContainsAny(str,nums) {
		count++
	}
	if strings.ContainsAny(str,letters) {
		count++
	}
	if strings.ContainsAny(str,upLetters) {
		count++
	}
	if strings.ContainsAny(str,others) {
		count++
	}
	if count <3 {
		return false
	}
	for i:=0;i<len(str)-2;i++ {
		if strings.Count(str,str[i:i+3]) >1 {
			return false
		}
	}
	return true
}


简单密码

package main

import (
    "bufio"
    "fmt"
    "os"
    //"strconv"
    //"strings"
)

func main() {
    inputReader := bufio.NewReader(os.Stdin)
    pwd, err := inputReader.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    iLen := len(pwd)
    var ch byte
    for i:=0; i<iLen - 1; i++ {
        if pwd[i] >= 'A' && pwd[i] <= 'Z' {
            ch = pwd[i] + 'a' - 'A' + 1
            if ch > 'z' {
                ch = 'a'
            }
            fmt.Printf("%c", ch)
        } else if pwd[i] >= 'a' && pwd[i] <= 'z' {
            if pwd[i] >= 'a' && pwd[i] <= 'c' {
                fmt.Printf("2")
            } else if pwd[i] >= 'd' && pwd[i] <= 'f' {
                fmt.Printf("3")
            } else if pwd[i] >= 'g' && pwd[i] <= 'i' {
                fmt.Printf("4")
            } else if pwd[i] >= 'j' && pwd[i] <= 'l' {
                fmt.Printf("5")
            } else if pwd[i] >= 'm' && pwd[i] <= 'o' {
                fmt.Printf("6")
            } else if pwd[i] >= 'p' && pwd[i] <= 's' {
                fmt.Printf("7")
            } else if pwd[i] >= 't' && pwd[i] <= 'v' {
                fmt.Printf("8")
            } else if pwd[i] >= 'w' && pwd[i] <= 'z' {
                fmt.Printf("9")
            }
        } else {
            fmt.Printf("%c", pwd[i])
        }
    }
    fmt.Printf("\n")
}


汽水瓶
package main

import "fmt"

func main() {
	var n int
	for {
		x,_ := fmt.Scan(&n)
		if x == 0 {
			break
		} else {
			count := 0
			remain := n
			for n/3 > 0 {
				remain = n % 3
				count += n / 3
				n = remain + n/3
			}

			if n == 2 {
				count++
			}
			if count > 0 {
				fmt.Println(count)
			}
		}
	}
}

删除字符串中出现次数最少的字符

package main

import (
    //     "bufio"
    //     "os"
    "fmt"
    "strings"
)

func main () {
    for {
        var input string
        fmt.Scanln(&input)
        if input == "" {
            return
        }

        inputArray := strings.Split(input, "")

        countMap := make(map[string]int, len(inputArray))
        for _, s := range inputArray {
            if _, exist := countMap[s]; !exist {
                countMap[s] = 1
            }else {
                countMap[s] += 1
            }
        }

        //fmt.Printf("%+v", countMap)

        min := countMap[inputArray[0]]
        var minArray []string
        minArray = make([]string, 0, len(inputArray))

        for k,v :=range countMap {
            //fmt.Println("min", min)
            if v < min {
                minArray = make([]string, 0, len(inputArray))
                minArray = append(minArray, k)
                min = v
            }else if v == min{
                //fmt.Println(k)
                minArray = append(minArray, k)
            }else {
                //min = v
            }
        }

        //fmt.Println(minArray)

        for _, k := range minArray {
            input = strings.ReplaceAll(input, k, "")
        }

        fmt.Println(input)
        }
}

合唱队

package main

import (
	"fmt"
)

var N, maxs = 0, [2]int{0,0}

func main() {
	for i := 0; i < 2; i++ {
		fmt.Scan(&N)
		maxs[i] = loopInput(N)
	}
	for _, v := range maxs {
		fmt.Println(v)
	}
}

func loopInput(N int) int {
	hights, inhights := make([]int, N), make([]int, N)
	asc, des := make([]int, N), make([]int, N)
	max := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&hights[i])
	}
	asc = getMax(N, hights)
	for i := range hights {
		inhights[i] = hights[N-1-i]
	}
	des = getMax(N, inhights)
	for i, v := range asc {
		if max < v+des[N-1-i]-1 {
			max = v + des[N-1-i] - 1
		}
	}
	return N - max
}

func getMax(N int, hights []int) []int {
	ace, ends := make([]int, N), make([]int, N)
	right, l, r, m := 0, 0, 0, 0
	ends[0], ace[0] = hights[0], 1
	for i := 1; i < len(hights); i++ {
		l = 0
		r = right
		for l <= r {
			m = (l + r) / 2
			if hights[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		if right < l {
			right = l
		}
		ends[l] = hights[i]
		ace[i] = l + 1
	}
	return ace
}


数据分类处理

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Ithems struct {
	index int
	value string
}

func main() {

	I := map[int]string{}
	inum := 0

	buf := bufio.NewReader(os.Stdin)

	for i := 0; i < 4; i++ {
		a, _, c := buf.ReadLine()
		if c != nil {
			return
		}
		if len(a) == 0 {
			break
		}
		I[inum] = string(a)
		inum++
	}
	for i := 0; i < inum; i = i + 2 {
		count(I[i], I[i+1])
	}

}

func count(I1, I2 string) {

	StrI := strings.Split(I1, " ")
	StrR := strings.Split(I2, " ")
	StrI, StrR = StrI[1:], StrR[1:]
	flag := map[string]bool{}
	tempR := []string{}
	for _, v := range StrR {
		if !flag[v] {
			tempR = append(tempR, v)
			flag[v] = true
		} else {
			continue
		}
	}
	// fmt.Println(tempR)
	nums := map[string]int{}
	resI := map[string][]Ithems{}
	for j := 0; j < len(tempR); j++ {
		for i := 0; i < len(StrI); i++ {
			if strings.Contains(StrI[i], tempR[j]) {
				// fmt.Println(tempR[j])
				nums[tempR[j]]++
				resI[tempR[j]] = append(resI[tempR[j]], Ithems{index: i, value: StrI[i]})
			}
		}
	}

	var add int = 0
	for i, v := range nums {
		if i == "" {
			delete(nums, i)
		} else {
			add = add + v
		}
	}
	AllNums := len(nums) * 2
	AllNums = AllNums + add*2

	fmt.Print(AllNums, " ")

	var key []int
	for k := range nums {
		kk, _ := strconv.Atoi(k)
		key = append(key, kk)
	}
	sort.Ints(key)

	for _, v := range key {
		m := strconv.Itoa(v)
		fmt.Print(v, " ", nums[m], " ")
		for _, v := range resI[m] {
			fmt.Print(v.index, " ", v.value, " ")
		}
	}
	fmt.Println()
}

字符串排序

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        sort(input.Text())
    }
}

func sort(s string) {
    symbolIndex := make([]int, 0)
    symbolValue := make([]byte, 0)
    //遍历字符判断，非字母就存index和value到两个数组里
    for i, v := range s {
        if v > 'z' || v > 'Z' && v < 'a' || v < 'A' {
            symbolIndex = append(symbolIndex, i)
            symbolValue = append(symbolValue, uint8(v))
        }
    }
    //fmt.Println(symbolIndex, symbolValue)
    // 遍历26个大小写字母，按a-z顺序存入数组char中
    char := make([]byte, 0)
    for i := 0; i < 26; i++ {
        for _, v := range s {
            value := uint8(v)
            if value == 'a'+uint8(i) || value == 'A'+uint8(i) {
                char = append(char, value)
            }
        }
    }
    //fmt.Println(char)
    //遍历符号位置信息存入结果数组中
    res := make([]byte, len(s))
    for i, v := range symbolValue {
        res[symbolIndex[i]] = v
    }
    //fmt.Println(res)
    //遍历数组若元素为0 则赋值为字符byte
    for _, c := range char {
        for i, v := range res {
            if v == 0 {
                res[i] = c
                break
            }
        }
    }
    fmt.Println(string(res))
}

查找兄弟单词

package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
    "sort"
    "strings"
    "strconv"
)

func main(){
    reader:=bufio.NewReader(os.Stdin)
    text,_,err:=reader.ReadLine()
    for err!=io.EOF{
        parse:=strings.Split(string(text)," ")
        end,_:=strconv.Atoi(parse[0])
        dict:=parse[1:end+1]
        var brother []string
        for _,word:=range dict{
            if isBrother(parse[len(parse)-2],word){
                brother=append(brother,word)
            }
        }
        sort.Strings(brother)
        n,_:=strconv.Atoi(parse[len(parse)-1])
        fmt.Println(len(brother))
        if len(brother)>=n{
            fmt.Println(brother[n-1])
        }
        text,_,err=reader.ReadLine()
    }
}

func isBrother(src,dst string)bool{
    if src==dst{
        return false
    }
    if len(src)!=len(dst){
        return false
    }
    var hash [26]int
    for i:=0;i<len(src);i++{
        hash[src[i]-'a']++
        hash[dst[i]-'a']--
    }
    for j:=0;j<26;j++{
        if hash[j]!=0{
            return false
        }
    }
    return true
}

素数伴侣
package main

import (
	"fmt"
)

func main() {
	for {
		var cnt int
		n, _ := fmt.Scan(&cnt)
		if n == 0 {
			break
		}
		var evens, odds []int
		for i := 0; i < cnt; i++ {
			var num int
			fmt.Scan(&num)
			if num%2 == 0 {
				evens = append(evens, num)
			} else {
				odds = append(odds, num)
			}
		}
		sum := 0
		bros := make([]int, len(evens))
		for _, odd := range odds {
		    used := make([]bool, len(evens))
			if find(odd, evens, bros, used) {
				sum++
			}
		}
		fmt.Println(sum)

	}
}

func find(odd int, evens []int, bros []int, used []bool) bool {
	for i := 0; i < len(evens); i++ {
		if isBro(evens[i], odd) && !used[i] {
			used[i] = true
			if bros[i]==0 || find(bros[i], evens, bros, used) {
				bros[i] = odd
				return true
			}
		}
	}
	return false
}

func isBro(a, b int) bool {
	return isPrime(a + b)
}

func isPrime(num int) bool {
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

字符串加解密

package main

import (
	"bufio"
	"os"
    "fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    i := 0
	for scanner.Scan() {
        if i % 2 == 0 {
            fmt.Println(encrypt(scanner.Text()))
        }else {
            fmt.Println(decrypt(scanner.Text()))
        }
        i++
	}
}

func encrypt(str string) string {
    data := []byte(str)
    for i := 0; i < len(data); i++{
        d := data[i]
        t := d
        if d >= 'a' && d < 'z'{
            t = d - 31
        }else if d == 'z'{
            t = 'A'
        }else if d == 'Z' {
            t = 'a'
        }else if d >= 'A' &&d < 'Z' {
            t = d + 33
        } else if d >= '0' &&d < '9'{
            t = d + 1
        } else if d == '9' {
            t = '0'
        }
        data[i] = t
    }
    return string(data)

}

func decrypt(str string) string{
    data := []byte(str)
    for i := 0; i < len(data); i++{
        d := data[i]
        t := d
        if d > 'a' && d <= 'z'{
            t = d - 33
        }else if d == 'a'{
            t = 'Z'
        }else if d == 'A' {
            t = 'z'
        }else if d > 'A' &&d <= 'Z' {
            t = d + 31
        } else if d > '0' &&d <= '9'{
            t = d - 1
        } else if d == '0' {
            t = '9'
        }
        data[i] = t
    }
    return string(data)
}

字符串合并处理

package main

import (
	"fmt"
	"io"
	"sort"
)

const helper1 string = "0123456789abcdefABCDEF"
const helper2 string = "084C2A6E195D3B7F5D3B7F"

func main() {

	var str1, str2 string
	for {
		if _, err := fmt.Scan(&str1, &str2); err == io.EOF {
			break
		}

		str := str1 + str2

		odds, evens := make([]byte, 0), make([]byte, 0)
		for i := 0; i < len(str); i++ {
			if i%2 == 0 {
				odds = append(odds, str[i])
			} else {
				evens = append(evens, str[i])
			}
		}

		sort.Slice(odds, func(i, j int) bool {
			return odds[i] < odds[j]
		})
		sort.Slice(evens, func(i, j int) bool {
			return evens[i] < evens[j]
		})

		ans := make([]byte, 0)

		for i := 0; i < len(str); i++ {
			if i%2 == 0 {
				ans = append(ans, fmtByte(odds[i/2]))
			} else {
				ans = append(ans, fmtByte(evens[i/2]))
			}
		}

		fmt.Println(string(ans))
	}
}

func fmtByte(num byte) byte {
	if num >= '0' && num <= '9' {
		num = num - '0'
		a := num&(1<<3)>>3 + (num&(1<<2))>>1 + (num&(1<<1))<<1 + (num&1)<<3
		d := fmt.Sprintf("%X", a)
		return d[0]
		//return fmt.Sprintf("%X", strconv.Itoa(int(num&(1<<3)>>3 + (num&(1<<2))>>1 + (num&(1<<1))<<1 + (num&1)<<3))[0])[0]
	} else if num >= 'a' && num <= 'f' {
		num = num - 'a' + 10
		a := num&(1<<3)>>3 + (num&(1<<2))>>1 + (num&(1<<1))<<1 + (num&1)<<3
		d := fmt.Sprintf("%X", a)
		return d[0]
	} else if num >= 'A' && num <= 'F' {
		num = num - 'A' + 10
		a := num&(1<<3)>>3 + (num&(1<<2))>>1 + (num&(1<<1))<<1 + (num&1)<<3
		d := fmt.Sprintf("%X", a)
		return d[0]
	} else {
		return num
	}
}

密码截取

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, _ := reader.ReadLine()
		if len(input) == 0 {
			break
		}
		data := string(input)
		arr := []rune(data)
		max := 1
		for i := 0; i < len(arr); i++ {
			left := i - 1
			right := i + 1
			count := 1
			for left >= 0 && right < len(arr) {
				if arr[left] != arr[i] && arr[right] != arr[i] {
					break
				}
				if arr[left] == arr[i] {
					left--
				}
				if arr[right] == arr[i] {
					right++
				}

			}
			count = right - left - 1
			if max < count {
				max = count
			}
			for left >= 0 && right < len(arr) {
				if arr[left] == arr[right] {
					left--
					right++
					count = count + 2
					if count > max {
						max = count
					}
				} else {
					break
				}

			}

		}
		fmt.Println(max)
	}
}

整数与IP地址间的转换

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var ipStr string
		var lonStr int64
		_, err1 := fmt.Scan(&ipStr)
		if err1 != nil {
			break
		}

		_, err2 := fmt.Scan(&lonStr)
		if err2 != nil {
			break
		}

		ipSlice := strings.Split(ipStr, ".")
		lastjieguo:=""
		for i := 0; i < len(ipSlice); i++ {
			ipInt64,_:=strconv.ParseInt(ipSlice[i],10,64)
			ipBinary:=strconv.FormatInt(ipInt64,2)
			result:=""
			if len(ipBinary)<8{
				for j:=0;j<8-len(ipBinary);j++{
					result+="0"
				}
			}
			lastjieguo=lastjieguo+result+ipBinary
		}
		a,_:=strconv.ParseInt(lastjieguo,2,64)
		fmt.Println(a)


		b:=strconv.FormatInt(lonStr,2)
		b1:=""
		if len(b)<32{
			for t:=0;t<32-len(b);t++{
				b1+="0"
			}

		}
		b1+=b

		//0-8
		//8-16
		//16-24
		//24-32
		result2:=""
		for w:=0;w<4;w++{
			wint,_:=strconv.ParseInt(b1[w*8:w*8+8],2,64)
			result2=result2+strconv.FormatInt(wint,10)+"."
		}
		fmt.Println(result2[0:len(result2)-1])
	}
}

图片整理

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str string
	for {
		_,err:= fmt.Scan(&str)
		if err != nil {
			return
		}
		sortImg(str)
	}

}

func sortImg(str string)  {
	var sli  []string
	for i:=0;i<len(str);i++ {
		sli = append(sli,string(str[i]) )
	}
	sort.Strings(sli)
	var sb strings.Builder
	for _,val := range sli {
		sb.WriteString(val)
	}
	fmt.Println(sb.String())
}

蛇形矩阵
package main
import(
    "fmt"
)

func main(){
    for{
        var size int
        if _, err :=fmt.Scanf("%d", &size); err !=nil || size<1 || size>100{
            return
        }
        arr1 := []int{}
        for i:=0; i<size; i++{
            arr := []int{}
            for j:=0; j<size-i; j++{
                if i==0{
                    if j==0{
                        arr1 = append(arr1,1)
                    }else{
                        arr1 = append(arr1,arr1[j-1]+j+1)
                    }
                }else{
                    arr = append(arr,arr1[j+1]-1)
                }
            }
            var data string
            if i==0{
                data = fmt.Sprint(arr1)
            }else{
                data = fmt.Sprint(arr)
                arr1 = arr
            }

            fmt.Println(data[1:len(data)-1])
        }
    }
}

字符串加密
package main

import(
"bufio"
"os"
    "io"
    "fmt"
)

func main(){
    reader :=bufio.NewReader(os.Stdin)
    for {
        keybyte,_,err:=reader.ReadLine()
        if err!=nil||err==io.EOF{
            return
        }
        strbyte,_,err:=reader.ReadLine()
         if err!=nil||err==io.EOF{
            return
        }
        letter_buf:=make([]byte,0)
        key_map:=make(map[byte]int,0)

        for _,v:=range keybyte{
            if _,ok:=key_map[v];!ok{
                key_map[v]=0
                letter_buf=append(letter_buf, v)
            }
        }

        for i:='a';i<='z';i++{
            if _,ok:=key_map[byte(i)];!ok{
                letter_buf=append(letter_buf, byte(i))
            }
        }

        letter_map:=make(map[byte]byte,0)
        letter:='a'
        for _,v:=range letter_buf{
            letter_map[byte(letter)]=v
            letter++
        }
        var buf []byte
        for _,v:=range strbyte{
            if v>='a'&&v<='z'{
                buf=append(buf, letter_map[v])
            }else if v>='A'&&v<='Z'{
                buf=append(buf, letter_map[v-32])
            }
        }
        fmt.Println(string(buf))
    }
}

统计每个月兔子的总数

package main

import (
    "fmt"
)

func main() {
    for {
        var n int
        num, _ := fmt.Scanln(&n)
        if num == 0{
            break
        }
        pre, now, next := 1, 0 ,0
        for i := 3; i <= n; i ++ {
            pre += now
            now = next
            next = pre
        }
        fmt.Println(pre + now + next)
    }
}

求小球落地5次后所经历的路程和第5次反弹的高度

package main
import (
"fmt"
)
func main(){

    var sum,high,tmp float64
    fmt.Scan(&high)
    tmp=high //初始高度算弹上来的算
    for i:=1;i<=5;i++{
       sum+=2*tmp //第i次落地时，经历的长度
        tmp=tmp/2 //第i+1次反弹的高度
    }
    fmt.Println(sum-high) //减去初始高度
    fmt.Println(tmp)
}

判断两个IP是否属于同一子网
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getBinaryNum(str string) error {
	splitStr := strings.Split(str, ".")
	if len(splitStr) < 4 {
		return errors.New("error")
	} else {
		for i := 0; i < len(splitStr); i++ {
			a, _ := strconv.ParseInt(splitStr[i], 10, 64)
			if a > 255 || a < 0 {
				return errors.New("error")
			}
		}
	}
	return nil
}

func panduanNetmask(str string) bool {
	splitStr := strings.Split(str, ".")
	result := ""
	for i := 0; i < len(splitStr); i++ {
		test, _ := strconv.ParseInt(splitStr[i], 10, 64)
		testStr := strconv.FormatInt(test, 2)
		testStrLen := len(testStr)
		if testStrLen < 8 {
			for j := 0; j < 8-testStrLen; j++ {
				result += "0"
			}
		}
		result = result+testStr
	}
	index0 := 0
	for index := 0; index < len(result); index++ {
		if result[index] == '0' {
			index0 = index
			break
		}
	}

	for a:=index0;a<len(result);a++{
		if result[a] == '1'{
			return true
		}
	}
	return false
}

func main() {


	scan := bufio.NewScanner(os.Stdin)
	for {
		scan.Scan()
		a := scan.Text()
		if a == "" {
			break
		}
		scan.Scan()
		b := scan.Text()
		if b == "" {
			break
		}
		scan.Scan()
		c := scan.Text()
		if c == "" {
			break
		}

		var err1,err2,err3 error
		err1 = getBinaryNum(a)
		err2 = getBinaryNum(b)
		err3 = getBinaryNum(c)
		ismast := panduanNetmask(a)
		if err1 != nil ||err2 != nil||err3 != nil||ismast{
			fmt.Println(1)
		} else {
			aSplit := strings.Split(a, ".")
			bSplit := strings.Split(b, ".")
			cSplit := strings.Split(c, ".")
			for i := 0; i < 4; i++ {
				aInt, _ := strconv.ParseInt(aSplit[i],10,64)
				bInt, _ := strconv.ParseInt(bSplit[i],10,64)
				cInt, _ := strconv.ParseInt(cSplit[i],10,64)
				if aInt&bInt != aInt&cInt {
					fmt.Println(2)
					break
				} else {
					if i==3{
						fmt.Println(0)
					}
					continue
				}
			}

		}
	}
}

统计字符
package main
import (
    "fmt"
    "bufio"
    "os"
)
func main(){
    input:=bufio.NewScanner(os.Stdin)
    for input.Scan(){
        s:=string(input.Bytes())
        m,n,p,q:=static(s)
        fmt.Println(m)
        fmt.Println(n)
        fmt.Println(p)
        fmt.Println(q)
    }
}

func static(s string)(int,int,int,int){
    m,n,p,q:=0,0,0,0
    for _,v:=range s{
        if v>='a'&&v<='z'||v>='A'&&v<='Z'{
            m++
        }else if v==' '{
            n++
        }else if v>='0'&&v<='9'{
            p++
        }else{
            q++
        }
    }
    return m,n,p,q
}

称砝码
// 称砝码

package main

import "fmt"

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		solve(n)
	}
}
func solve(n int) {
	m := make([]int, n)
	c := make([]int, n)
	sum := 0
	for i, _ := range m {
		fmt.Scan(&m[i])
	}
	for i, _ := range c {
		fmt.Scan(&c[i])
		sum += c[i] * m[i]
	}
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < c[i]; j++ {
			for k := sum; k >= m[i]; k-- {
				if dp[k-m[i]] {
					dp[k] = true
				}
			}
		}
	}
	cnt := 0
	for i := range dp {
		if dp[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}

学英语
//学英语

package main

import (
	"fmt"
	"math"
)

func zeroToNineteen(num int64) {
	exps := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	fmt.Print(exps[num-1])
}

func twentyToNinetyNine(num int64) {
	if num < 20 {
		zeroToNineteen(num)
		return
	}
	a, b := num/10, num%10
	exps := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	fmt.Print(exps[a-2])
	if b != 0 {
		fmt.Print(" ")
		zeroToNineteen(b)
	}
}

func oneHundredTo999(num int64) {
	if num < 100 {
		twentyToNinetyNine(num)
		return
	}
	exps := " hundred"
	a, b := num/100, num%100
	zeroToNineteen(a)
	fmt.Print(exps)
	if b != 0 {
		fmt.Print(" and ")
		twentyToNinetyNine(b)
	}
}

func oneThousandTo999999(num int64) {
	if num < 1000 {
		oneHundredTo999(num)
		return
	}
	exps := " thousand "
	a, b := num/1000, num%1000
	oneHundredTo999(a)
	fmt.Print(exps)
	if b != 0 {
		oneHundredTo999(b)
	}
}

func oneMillionTo999999999(num int64) {
	if num < 1000000 {
		oneThousandTo999999(num)
		return
	}
	exps := " million "
	a, b := num/1000000, num%1000000
	oneHundredTo999(a)
	fmt.Print(exps)
	if b != 0 {
		oneThousandTo999999(b)
	}
}

func oneBillionToInt32Max(num int64) {
	if num < 1000000000 {
		oneMillionTo999999999(num)
		return
	}
	exps := " billion "
	a, b := num/1000000000, num%1000000000
	oneHundredTo999(a)
	fmt.Print(exps)
	if b != 0 {
		oneMillionTo999999999(b)
	}
}

func numToEng(num int64) {
	if num == 0 {
		fmt.Println("zero")
		return
	}
	if num < 0 {
		fmt.Print("negative ")
	}
	num = int64(math.Abs(float64(num)))

	oneBillionToInt32Max(num)
	fmt.Println()

}

func main() {
	for {
		var num int64
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		numToEng(num)
	}
}

迷宫问题

package main

import (
    "fmt"
)

func main() {
    var a int
    var row int
    var col int
    for {
        n,_:=fmt.Scan(&row)
        if n==0{
            break
        }
      n,_=fmt.Scan(&col)
        if n==0{
            break
        }
      mp:=make([][]int,row)
    for i:=0;i<row;i++{
        mp[i]=make([]int, col)
    }

      for i:=0;i<row;i++{
          for j:=0;j<col;j++{
              n,_:=fmt.Scan(&a)
              if n == 0 {
            break
              }else{
               mp[i][j]=a
              }
          }
    }
      check(mp,0,0)
     printPath(mp,0,0)
    }

}

func check(mp [][]int,x int,y int)int {
    if x==len(mp)-1&&y==len(mp[x])-1{
        return 1
    }
    if x<len(mp)&&y<len(mp[len(mp)-1])&&mp[x][y]==0{
        if check(mp,x+1,y)==-1 &&check(mp,x,y+1)==-1{
           mp[x][y]=1
            return -1
        }else{
            return 1
        }

    }else{
        return -1
    }
}
func printPath(mp [][]int,x int,y int){
    if x==len(mp)-1&&y==len(mp[x])-1{
        fmt.Printf("(%d,%d)\n",x,y)
        return
    }
     if x<len(mp)&&y<len(mp[len(mp)-1])&&mp[x][y]==0{
    fmt.Printf("(%d,%d)\n",x,y)
      printPath(mp,x+1,y)
      printPath(mp,x,y+1)

     }
}

Sudoku
package main

import (
    "fmt"
)

func main() {
	for {
		input := make([][]int, 9)
		var temp int
		var count int
		for i := range input {
			input[i] = make([]int, 9)
			for j := range input[i] {
				_, err := fmt.Scan(&temp)
				if err != nil {
					break
				}
				if temp == 0 {
					count++
				}
				input[i][j] = temp
			}
		}
		dfs(input, count)
		for i:= range input {
			for j := range input[i] {
				fmt.Printf("%d ", input[i][j])
			}
			fmt.Println()
		}
        break
	}

}

func dfs(input [][]int, count int) int {
	if count == 0 {
		return 0
	}
	for i:=0;i<len(input);i++ {
		for j:=0;j<len(input[i]);j++ {
			if input[i][j] == 0 {
				for x:=1;x<=10;x++ {
					if x == 10 {
						return count
					}
					if !isValid(input, i, j, x) {
						continue
					}
					input[i][j] = x
					if dfs(input, count - 1) == 0 {
						return 0
					}
					input[i][j] = 0
				}
			}
		}
	}
	return count
}

func isValid(input [][]int, i, j, v int ) bool {
	for x := 0;x<9;x++ {
		if input[x][j] == v {
			return false
		}
	}
	for y := 0;y<9; y++ {
		if input[i][y] == v {
			return false
		}
	}
	tx := i / 3
	ty := j / 3
	for x:=0;x<3;x++ {
		for y:=0;y<3;y++ {
			if input[tx*3+x][ty*3+y] == v {
				return false
			}
		}
	}
	return true
}

名字的漂亮度

package main

import(
"fmt"
    "io"
    "bufio"
    "strconv"
    "sort"
    "os"
)
/*
1、最大"漂亮度"==字母数相同次数最多的 漂亮度最大

*/
func main(){
    var num int
    reader:=bufio.NewReader(os.Stdin)
    for {
        numbyte,_,err:=reader.ReadLine()
        if err!=nil||err==io.EOF{
            return
        }
        num,_=strconv.Atoi(string(numbyte))
        for i:=0;i<num;i++{
            m:=make(map[byte]int,0)
            strbyte,_,_:=reader.ReadLine()
            //统计相同字母的个数
            for _,v:=range strbyte{
                if _,ok:=m[v];ok{
                    m[v]+=1
                }else{
                    m[v]=1
                }
            }

            //根据字母相同个数排序+不需要保留字母了
            count_sort:=make([]int,0)
            for _,v:=range m{
                count_sort=append(count_sort, v)
            }
            sort.Slice(count_sort,func (i,j int )bool{
                return count_sort[i]>count_sort[j]
            })
            sum:=0
            for i,v:=range count_sort{
                sum+=v*(26-i)
            }

            fmt.Println(sum)

        }
    }
}

截取字符串
package main

import(
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main(){
    scanner:=bufio.NewScanner(os.Stdin)
    for{
        scanner.Scan()
    s:=scanner.Text()
        if len(s)<=0{
            break
        }
    scanner.Scan()
    n,_:=strconv.Atoi(scanner.Text())

    for i:=0;i<len(s);i++{
        if i<n{
            fmt.Print(string(s[i]))
        }else{
            break
        }
    }
    fmt.Println()
    }
}

从单向链表中删除指定值的节点
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type ListNode struct {
    Key  int
    Next *ListNode
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        s1, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        s1 = s1[:len(s1)-1]
        sList := strings.Split(s1, " ")

        num, _ := strconv.Atoi(sList[0])
        first, _ := strconv.Atoi(sList[1])

        head := &ListNode{
            Key: first,
        }

        k := 2
        for i := 0; i < num-1; i++ {
            key1, _ := strconv.Atoi(sList[k])
            key2, _ := strconv.Atoi(sList[k+1])

            for link := head; link != nil; link = link.Next {
                if link.Key == key2 {
                    newNode := &ListNode{
                        Key:  key1,
                        Next: link.Next,
                    }
                    link.Next = newNode
                    break
                }
            }
            k = k + 2
        }

        d, _ := strconv.Atoi(sList[k])

        if head != nil && head.Key == d {
            head = head.Next
        } else {
            for link := head; link != nil && link.Next != nil; link = link.Next {
                if link.Next.Key == d {
                    link.Next = link.Next.Next
                    break
                }
            }
        }

        for link := head; link != nil; link = link.Next {
            fmt.Printf("%d ", link.Key)
        }
        fmt.Printf("\n")

    }

}

多线程
package main

import(
    "fmt"
)

func main(){
    var n int
    for {
        num,_ :=fmt.Scan(&n)
        if num == 0 {
            break
        }
        for i:=0;i<n;i++{
            fmt.Printf("ABCD")
        }
        fmt.Printf("\n")
    }

}

四则运算
package main

import(
    "fmt"
    "strings"
    "strconv"
)

func main () {
    for {
        var s string
        n, _ := fmt.Scanln(&s)
        if n == 0 {
            break
        }
        s = strings.ReplaceAll(s, "{", "(")
        s = strings.ReplaceAll(s, "}", ")")
        s = strings.ReplaceAll(s, "[", "(")
        s = strings.ReplaceAll(s, "]", ")")


                for i := 0; i < len(s); i ++ {
            if s[i] == '-' {
                if i == 0 || ((s[i - 1] < '0' || s[i -1] > '9') && s[i-1]!=')') {
                    var j int
                    for j = 1; i + j < len(s); j ++ {
                        if s[i + j] < '0' || s[i + j] > '9' {
                            break
                        }
                    }
                    s = s[:i] + "(0-" + s[i + 1:i + j] + ")" +s[i+j:]
                }
            }
        }
        //fmt.Print(s)
        arr := make([]string, 0)
        i := 0
        for i< len(s) {
            var temp string
            for i < len(s) && s[i] <= '9' && s[i] >= '0' {
                temp += string(s[i])
                i ++
            }
            if temp != "" {
                arr = append(arr, temp)
            }
            i ++
        }
        //fmt.Print(arr, len(arr))
        iarr := 0
        op := make([]byte, 100)
        iop := 0
        nums := make([]string, 100)
        inums := 0
//        flag := false
        for i = 0; i < len(s); i ++ {
            //fmt.Print(i, " ")
            if s[i] == '(' || s[i] == '*' || s[i] == '/' {
                op[iop] = s[i]
                iop ++
            } else if s[i] == '+' || s[i] == '-' {
                    for iop > 0 && op[iop - 1] != '(' {
                        nums[inums] = string(op[iop - 1])
                        inums ++
                        op[iop - 1] = s[i]
                        iop --
                    }
                op[iop] = s[i]
                iop ++
                    //fmt.Print("+++++++++",op)
            } else if s[i] == ')' {
                for op[iop - 1] != '(' {
                    nums[inums] = string(op[iop - 1])
                    inums ++
                    iop --
                }
                iop --
            } else {
                for (i < len(s) && s[i] >= '0' && s[i] <= '9') {
                    i ++
                }
                nums[inums] = arr[iarr]
                inums ++
                iarr ++
                i --
            }
        }
        //fmt.Print(iop)
        for iop > 0 {
            nums[inums] = string(op[iop - 1])
            //fmt.Print(string(op[iop - 1]))
            iop --
            inums ++
        }
        cacl := make([]int, 100)
        icacl := 0
        for i := 0; i < inums; i ++ {
            //fmt.Print(nums[i], " ")
            if nums[i] != "+" && nums[i] != "-" && nums[i] != "*" && nums[i] != "/" {
                temp, _ := strconv.Atoi(nums[i])
                cacl[icacl] = temp
                icacl ++
            } else {
                res := 0
                if nums[i] == "+" {
                    res = cacl[icacl - 2] + cacl[icacl - 1]
                } else if nums[i] == "-" {
                    res = cacl[icacl - 2] - cacl[icacl - 1]
                } else if nums[i] == "*" {
                    res = cacl[icacl - 2] * cacl[icacl - 1]
                } else {
                    res = cacl[icacl - 2] / cacl[icacl - 1]
                }
                //fmt.Print(res, " ")
                cacl[icacl - 2] = res
                icacl --
            }
        }
        fmt.Println(cacl[0])
    }
}

输出单向链表中倒数第k个结点
package main
import     "fmt"
type node struct {
    key int
    next *node
}
func main() {
    var (
        n int
        last int
    )
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }

        root := &node{}
        fmt.Scan(&root.key)

        ne := root
        for i := 0; i < n-1; i++ {
            ne.next = &node{}
            fmt.Scan(&ne.next.key)
            ne = ne.next
        }
        fmt.Scan(&last)

        if last == 0 {
            fmt.Println(0)
            continue
        }

        ne = root
        for i := 0; i < n-last; i++ {
            ne = ne.next
        }
        fmt.Println(ne.key)
        // fmt.Println(root.key, root.next.key)
    }
}

计算字符串的距离
package main

import (
    "fmt"
)

// f[i][j] =

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}

func leastDistance(s1, s2 string) int {
    len1, len2 := len(s1), len(s2)
    dp := make([][]int, len1+1)
    for i:=0; i<=len1; i++ {
        dp[i] = make([]int, len(s2)+1)
    }

    for i:=0; i<=len2; i++ {
        dp[0][i] = i
    }

    for i:=0; i<=len1; i++ {
        dp[i][0] = i
    }

    for i:=1; i<=len1; i++ {
        for j:=1; j<=len2; j++ {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1)
                dp[i][j] = min(dp[i][j], dp[i-1][j-1] + 1)
            }
        }
    }

    return dp[len1][len2]
}

func main() {
    for {
        var s1, s2 string
        fmt.Scanf("%s", &s1)
        fmt.Scanf("%s", &s2)

        if s1 == "" || s2 == "" {
            break
        }

        fmt.Printf("%d\n", leastDistance(s1, s2))
    }

}

字符逆序
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(str string) string {
	newStr := ""
	strSplit := strings.Split(str, "")
	for i := len(strSplit) - 1; i >= 0; i-- {
		newStr = newStr + strSplit[i]
	}
	return newStr
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	newStr := reverse(str)
	fmt.Println(newStr)
}

求解立方根

package main
import(
    "fmt"
)

func main(){
    var num float64
    flag:=false
    _,err:=fmt.Scan(&num)
    if err !=nil{
        return
    }
    if num<0{
       flag=true
       num=-num
    }

    result:=getCubicValue(num)
    if flag {
        fmt.Printf("%.1f\n",-result)
    }else{
        fmt.Printf("%.1f\n",result)
    }
}

func getCubicValue(num float64)float64{
    min:=0.0
    max:=num
    if num <1{
        max=1
    }

    for max-min >0.00000001{
        mid:=(max+min)/2
        if mid*mid*mid> num{
            max=mid
        }else if mid*mid*mid < num{
            min=mid
        } else{
            return mid
        }
    }
     return max
}

求最小公倍数

package main
import (
"fmt"
)
func main(){
    a:=0
    b:=0
    fmt.Scanf("%d",&a)
    fmt.Scanf("%d",&b)
    if a>b{
        a,b=b,a
    }
    ma:=a
    for {
        if a%ma==0 && b%ma==0{
            break
        }
        ma--
    }
    re:=a*b/ma
    fmt.Println(re)
}

单词倒排
package main

import (
    "os"
    "fmt"
    "bufio"
)

// A~Z 65~90
// a~z 97~122
// ' ' 32

func main() {
    data, _, _ := bufio.NewReader(os.Stdin).ReadLine()
    var tmp = 0
    for index := len(data) -1; index >= 0; index-- {
        if data[index] < 65 || data[index] > 122 || data[index] < 97 && data[index] > 90 {
            fmt.Printf("%s ", string(data[index + 1:index + tmp + 1]))
            tmp = 0
        } else {
            tmp++
            if index == 0 {
                fmt.Printf("%s", string(data[:tmp]))
            }
        }
    }

}



















































































































