

### 字符串或者数字倒叙
``` golang
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	for {
		var input int
		if _, err := fmt.Scan(&input); err != nil {
			return
		}
		inStr := strconv.Itoa(input)
		inSplit := strings.Split(inStr, "")

		reStr := ""
		//for _, k := range inputStr {
		//	revert = string(k) + revert
		//}
		for i := len(inSplit) - 1; i >= 0; i-- {
			reStr = reStr + inSplit[i]
		}

		fmt.Println(reStr)
	}
}
```

### 汽水瓶



### 等差数列求和
``` go
package main

import "fmt"

func main() {
	// 等差数列 2，5，8，11，14。。。
	// 求前n项的和
	// 第n个数为 3*n-1

	for {
		var n int
		if _, err := fmt.Scan(&n); err != nil {
			return
		}
		sum := 3*5*(n+1)*n/10 - n
		fmt.Println(sum)
	}
}
```

### 最小公倍数
``` go
package main

import "fmt"

func main() {
	//最小公倍数
	for {
		var a, b int
		if _, err := fmt.Scan(&a, &b); err != nil {
			return
		}

		if a > b {
			a, b = b, a
		}

		c := a
		for {
			if c%a == 0 && c%b == 0 {
				fmt.Println(c)
				break
			}
			c++
		}
	}
}
```

### 单词倒排



### 找出字符串中第一个只出现一次的字符
``` go
package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var input string
		if _, err := fmt.Scan(&input); err != nil {
			return
		}

		inStr := strings.Split(input, "")

		st := make(map[string]int)
		for _, v := range inStr {
			_, exits := st[v]
			if exits {
				st[v]++
			} else {
				st[v] = 1
			}
		}

		var t string
		for _, v := range inStr {
			iv, exits := st[v]
			if exits && iv == 1 {
				t = v
				break
			}
		}

		if len(t) == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(t)
		}
	}
}
```

### 查找输入整数二进制中1的个数
``` go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	for {
		var input int
		if _, err := fmt.Scan(&input); err != nil {
			return
		}

		// 将十进制数字转化为二进制字符串
		b := ""
		for input > 0 {
			t := input % 2
			b = strconv.Itoa(t) + b
			input = input / 2
		}

		nums := 0
		for _, v := range strings.Split(b, "") {
			if v == "1" {
				nums++
			}
		}
		fmt.Println(nums)
	}
}
```

### 查找两个字符串a,b中的最长公共子串
``` go
package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		s1, _ := reader.ReadString('\n')
		s2, _ := reader.ReadString('\n')

		if len(s1) > len(s2) {
			s1, s2 = s2, s1
		}
	}
	//TODO
}

```

### 输入n个整数，输出其中最小的k个
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var n, k int
	for {
		if scanner.Scan() {
			n, _ = strconv.Atoi(scanner.Text())
		} else {
			return
		}

		if scanner.Scan() {
			k, _ = strconv.Atoi(scanner.Text())
		}

		nums := make([]int, 0)
		for i := 0; i < n; i++ {
			if scanner.Scan() {
				item, _ := strconv.Atoi(scanner.Text())
				nums = append(nums, item)
			}
		}

		sort.Ints(nums)
		for i := 0; i < k; i++ {
			fmt.Printf("%d ", nums[i])
		}
	}
}
```

### 查找组成一个偶数最接近的两个素数
``` go
package main

import "fmt"

// 用于判断是否为素数
func isP(n int) bool {
	if n == 1 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for {
		var input int
		if _, err := fmt.Scan(&input); err != nil {
			return
		}
		if input%2 != 0 || input < 2 {
			return
		}

		mid := input / 2
		a := mid - 1
		b := mid + 1
		for a > 0 {
			if isP(a) && isP(b) {
				fmt.Println(a)
				fmt.Println(b)
				break
			}
			a--
			b++

		}
	}
}
```

### 字符串排序
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var n int
	for {
		if scanner.Scan() {
			n, _ = strconv.Atoi(scanner.Text())
		} else {
			return
		}

		strs := make([]string, n)
		for i := 0; i < n; i++ {
			if scanner.Scan() {
				strs[i] = scanner.Text()
			}
		}

		sort.Strings(strs)
		for i := 0; i < n; i++ {
			fmt.Println(strs[i])
		}
	}
}
```

### 提取不重复的整数
``` go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type empty struct{}

func main() {
	for {
		var input int
		if _, err := fmt.Scan(&input); err != nil {
			return
		}

		inStr := strconv.Itoa(input)
		inSlice := strings.Split(inStr, "")

		ouStr := ""
		containd := make(map[string]empty)
		for i := len(inSlice) - 1; i >= 0; i-- {
			t := inSlice[i]
			_, exits := containd[t]
			if !exits {
				ouStr = ouStr + t
				containd[t] = empty{}
			}
		}

		ouInt, _ := strconv.Atoi(ouStr)
		fmt.Println(ouInt)
	}
}
```

### 字符统计
``` go
package main

import (
	"fmt"
	"sort"
	"strings"
)

type items []item

type item struct {
	key   string
	value int
}

func (its items) Len() int {
	return len(its)
}

func (its items) Less(i, j int) bool {
	if its[i].value == its[j].value {
		return its[i].key < its[j].key
	}
	return its[i].value > its[j].value
}

func (its items) Swap(i, j int) {
	its[i], its[j] = its[j], its[i]
}

func NewItems(sts map[string]int) items {
	is := make([]item, 0)
	for k, v := range sts {
		is = append(is, item{
			key:   k,
			value: v,
		})
	}

	return is
}

func main() {
	for {
		var input string
		if _, err := fmt.Scan(&input); err != nil {
			return
		}
		inSlice := strings.Split(input, "")

		st := make(map[string]int)
		for _, str := range inSlice {
			_, exits := st[str]
			if exits {
				st[str]++
			} else {
				st[str] = 1
			}
		}

		nis := NewItems(st)
		sort.Sort(nis)

		var res string
		for _, i := range nis {
			res = res + i.key
		}

		fmt.Println(res)
	}
}
```

### 输入整型数组和排序标识，对其元素按照升序或降序进行排序

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var n int
		n, _ = strconv.Atoi(scanner.Text())

		sm := make([]int, 0)
		if scanner.Scan() {
			sSlice := strings.Split(scanner.Text(), " ")
			for i := 0; i < n; i++ {
				strInt, _ := strconv.Atoi(sSlice[i])
				sm = append(sm, strInt)
			}
		}
		sort.Ints(sm)

		var f int
		if scanner.Scan() {
			f, _ = strconv.Atoi(scanner.Text())
		}

		if f == 0 {
			for _, a := range sm {
				fmt.Printf("%d ", a)
			}
		}
		if f == 1 {
			for i := len(sm) - 1; i >= 0; i-- {
				fmt.Printf("%d ", sm[i])
			}
		}
		fmt.Println()
	}
}
```

### 数组分组

``` go
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		var n int
		n, _ = strconv.Atoi(scanner.Text())

		intSlice := make([]int, 0)
		if scanner.Scan() {
			strSlice := strings.Split(scanner.Text(), " ")
			for i := 0; i < n; i++ {
				strInt, _ := strconv.Atoi(strSlice[i])
				intSlice = append(intSlice, strInt)
			}
		}

		fSlice := make([]int, 0)
		tSlice := make([]int, 0)
		for _, is := range intSlice {
			if is%5 == 0 {
				fSlice = append(fSlice, is)
			} else if is%3 == 0 {
				tSlice = append(tSlice, is)
			}
		}

		// TODO
	}
}
```

### 记负均正
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var n int
	for {
		if scanner.Scan() {
			n, _ = strconv.Atoi(scanner.Text())
		} else {
			return
		}

		intSlice := make([]int, 0)
		for i := 0; i < n; i++ {
			if scanner.Scan() {
				strInt, _ := strconv.Atoi(scanner.Text())
				intSlice = append(intSlice, strInt)
			}
		}

		var sum float64
		count := 0
		for _, i := range intSlice {
			if i < 0 {
				count++
				continue
			}
			sum = sum + float64(i)
		}
		positeNum := float64(len(intSlice) - count)

		fmt.Printf("%d %.1f", count, sum/positeNum)
	}
}
```

### 记负均正II

```go

```

### 表示数字
```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	return true
}

func main() {
	for {
		var input string
		if _, err := fmt.Scan(&input); err != nil {
			return
		}
		inStr := strings.Split(input, "")

		var index bool
		var tmp string
		var ss string

		for _, is := range inStr {
			if isNumber(is) {
				if !index {
					index = true
				}
				tmp = tmp + is
			} else {
				index = false
				if len(tmp) != 0 {
					ss = ss + "*" + tmp + "*"
					tmp = ""
				}
				ss = ss + is
			}
		}
		ss = ss + "*" + tmp + "*"
		fmt.Println(ss)
	}
}
```

### 统计大写字母个数
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inStr := scanner.Text()

		var inub int
		// 大写字符(A-Z)的byte 为 65 - 90
		for index := range inStr {
			if inStr[index] >= 65 && inStr[index] <= 90 {
				inub++
			}
		}
		fmt.Println(inub)
	}
}
```

### 字符串最后一个单词的长度
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		s := strs[len(strs)-1]

		fmt.Println(len(s))
	}
}
```

### 字符串分隔
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
)

const L = 8

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			continue
		}

		var tmp string
		for i := 0; i < L-len(str)%L; i++ {
			tmp = tmp + "0"
		}

		if len(tmp) != 0 {
			str = str + tmp
		}

		for i := 0; i < len(str)/L; i++ {
			fmt.Println(str[i*L : (i+1)*L])
		}
	}
}
```

### 进制转换
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const L = 16

var CMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}

func Converted(index int, s string) int {
	num := CMap[s]
	tmp := 1
	for i := 0; i < index-1; i++ {
		tmp = tmp * L
	}

	return num * tmp
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		intStr := scanner.Text()
		if len(intStr) == 0 {
			continue
		}

		intSlice := strings.Split(intStr[2:], "")

		var s int
		for i := len(intSlice) - 1; i >= 0; i-- {
			s = s + Converted(len(intSlice)-i, intSlice[i])
		}
		fmt.Println(s)
	}
}
```

### 进制转换2
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		intStr := scanner.Text()
		if len(intStr) == 0 {
			continue
		}

		i10, _ := strconv.ParseInt(intStr[2:], 16, 32)
		fmt.Println(i10)
	}
}
```

### 质数因子
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPrime(it int) bool {
	if it == 1 {
		return true
	}

	for i := 2; i < it; i++ {
		if it%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var input int
		intStr := scanner.Text()
		if len(intStr) == 0 {
			continue
		}

		input, _ = strconv.Atoi(intStr)
		for !isPrime(input) {
			for i := 2; i < input; i++ {
				if isPrime(i) && input%i == 0 {
					input = input / i
					fmt.Printf("%d ", i)
					break
				}
			}
		}

		fmt.Printf("%d ", input)
	}
}
```

### 在字符串中找出连续最长的数字串
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 用于判断是否是数字字符串
func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

type items []item

type item struct {
	str   string
	lens  int
	index int
}

func (it items) Len() int      { return len(it) }
func (it items) Swap(i, j int) { it[i], it[j] = it[j], it[i] }

func (it items) Less(i, j int) bool {
	if it[i].lens == it[j].lens {
		return it[i].index < it[j].index
	}
	return it[i].lens > it[j].lens
}

func NewItems(s []string) items {
	its := make([]item, 0)
	for k, v := range s {
		its = append(its, item{
			str:   v,
			lens:  len(v),
			index: k,
		})
	}

	return its
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		intStr := scanner.Text()
		if len(intStr) == 0 {
			continue
		}

		var cursor bool
		var tStr string
		rSlice := make([]string, 0)

		for _, st := range strings.Split(intStr, "") {
			if isNumber(st) {
				if !cursor {
					cursor = true
				}
				tStr += st
			} else {
				if cursor {
					rSlice = append(rSlice, tStr)
					cursor = false
					tStr = ""
				}
			}
		}
		// 特殊处理字符串最后是数字串的场景
		rSlice = append(rSlice, tStr)

		its := NewItems(rSlice)
		sort.Sort(its)

		t := its[0].str
		n := its[0].lens

		for i := 1; i < len(its)-1; i++ {
			if its[i].lens == n {
				t = t + its[i].str
			}
		}
		fmt.Printf("%s,%d", t, n)
		// 格式要求不一致的时候，可以多打印一行
		fmt.Println()
	}
}
```

### 在字符串中找出连续最长的数字串2
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		intStr := scanner.Text()
		if len(intStr) == 0 {
			continue
		}
		// 正则匹配连续的数字
		re := regexp.MustCompile(`\d+`)
		result := re.FindAllString(intStr, -1)

		maxLen := 0
		for _, v := range result {
			if len(v) > maxLen {
				maxLen = len(v)
			}
		}

		var ret string
		for _, v := range result {
			if len(v) == maxLen {
				ret = ret + v
			}
		}
		fmt.Printf("%s,%d", ret, maxLen)
		fmt.Println()
	}
}
```

### 统计字符
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNum(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	return true
}

func isStr(s string) bool {
	// 小写字母
	if s[0] >= 65 && s[0] <= 90 {
		return true
	}
	// 大写字母
	if s[0] >= 97 && s[0] <= 122 {
		return true
	}

	return false
}

func isSpace(s string) bool {
	if s[0] == 32 {
		return true
	}
	return false
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strSlice := strings.Split(scanner.Text(), "")
		var strNum, spaceNum, nuNum, othNum int
		for _, s := range strSlice {
			if isStr(s) {
				strNum++
			} else if isSpace(s) {
				spaceNum++
			} else if isNum(s) {
				nuNum++
			} else {
				othNum++
			}
		}
		fmt.Println(strNum)
		fmt.Println(spaceNum)
		fmt.Println(nuNum)
		fmt.Println(othNum)
	}
}
```

### 自守数
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isZ(i int) bool {
	if strings.HasSuffix(strconv.Itoa(i*i), strconv.Itoa(i)) {
		return true
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n int
		n, _ = strconv.Atoi(scanner.Text())

		var count int
		for i := 0; i <= n; i++ {
			if isZ(i) {
				count++
			}
		}
		fmt.Println(count)
	}
}
```

### 名字的漂亮度
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type items []item

func (p items) Len() int           { return len(p) }
func (p items) Less(i, j int) bool { return p[i].num > p[j].num }
func (p items) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func NewItems(m map[string]int) items {
	nItems := make([]item, 0)
	for k, v := range m {
		nItems = append(nItems, item{
			str: k,
			num: v,
		})
	}
	return nItems
}

type item struct {
	str string
	num int
}

func calculateName(name string) int {
	nMap := make(map[string]int)
	for _, n := range strings.Split(name, "") {
		_, extes := nMap[n]
		if extes {
			nMap[n]++
		} else {
			nMap[n] = 1
		}
	}

	nItems := NewItems(nMap)
	sort.Sort(nItems)

	sum := 0
	for i, j := 0, 26; i < len(nItems); i, j = i+1, j-1 {
		sum = sum + nItems[i].num*j
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n int
		n, _ = strconv.Atoi(scanner.Text())

		names := make([]string, 0)
		for i := 0; i < n; i++ {
			if scanner.Scan() {
				// 忽略大小写
				names = append(names, strings.ToLower(scanner.Text()))
			}
		}

		for _, name := range names {
			fmt.Println(calculateName(name))
		}
	}
}
```

### 输出单向链表中倒数第k个结点
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	int      int
	ListNode *ListNode
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n int
		n, _ = strconv.Atoi(scanner.Text())

		nodes := make([]int, 0)
		if scanner.Scan() {
			vals := strings.Split(scanner.Text(), " ")
			for i := len(vals) - 1; i >= 0; i-- {
				valInt, _ := strconv.Atoi(vals[i])
				nodes = append(nodes, valInt)
			}
		}

		var k int
		if scanner.Scan() {
			k, _ = strconv.Atoi(scanner.Text())
		}

		// 构造链表
		fmt.Println(n, k)

		listNode := new(ListNode)
		for k1, v := range nodes {
			listNode.int = v

		}

		fmt.Println(nodes)
	}
}
```

### 输出单向链表中倒数第k个结点
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n int
		n, _ = strconv.Atoi(scanner.Text())
		if n == 0 {
			continue
		}

		// 构造单向链表
		head := &Node{}
		cur := head
		if scanner.Scan() {
			vals := strings.Split(scanner.Text(), " ")
			for i := 0; i < n; i++ {
				val, _ := strconv.Atoi(vals[i])
				tmp := &Node{val: val}
				cur.next = tmp
				cur = cur.next
			}
		}

		var k int
		if scanner.Scan() {
			k, _ = strconv.Atoi(scanner.Text())
		}

		// 长度已知，不需要通过双指针完成倒数第k个节点
		c := head.next
		for i := 1; i < n-k+1; i++ {
			c = c.next
		}

		fmt.Println(c.val)
	}
}
```

### 从单向链表中删除指定值的节点
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nStr := scanner.Text()
		if len(nStr) == 0 {
			return
		}

		nodes := strings.Split(nStr, " ")
		// n 个节点
		n, _ := strconv.Atoi(nodes[0])
		// 到节点的值
		headVal, _ := strconv.Atoi(nodes[1])
		// 需要删除的值
		delVal, _ := strconv.Atoi(nodes[len(nodes)-1])

		listNode := nodes[2 : len(nodes)-1]

		head := &Node{val: headVal}
		cur := head

		//for i := 0; i < n; i++ {
		//	val, _ := strconv.Atoi(nodes[i])
		//	tmp := &Node{val: val}
		//	cur.next = tmp
		//	cur = cur.next
		//}
		//
		//c := head.next
		//for i := 0; i < n; i++ {
		//	fmt.Println(c.val)
		//	c = c.next
		//}

	}
}
```


### 简单错误记录

### 识别有效的IP地址和掩码并进行分类统计

### 判断两个IP是否属于同一子网
``` go
// 判断两个ip是否属于同一子网
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var netmask string
		var ip1, ip2 string
		_, err := fmt.Scan(&netmask)
		if err != nil {
			break
		}

		_, err = fmt.Scan(&ip1)
		if err != nil {
			break
		}

		_, err = fmt.Scan(&ip2)
		if err != nil {
			break
		}

		calc(netmask, ip1, ip2)
	}
}

func calc(netmask, ip1, ip2 string) {
	b1, nsArr := isValidIp(netmask)
	if !b1 {
		fmt.Println(1)
		return
	}

	b1 = isValidNetmask(nsArr)
	if !b1 {
		fmt.Println(1)
		return
	}

	b2, ip1Arr := isValidIp(ip1)
	if !b2 {
		fmt.Println(1)
		return
	}

	b3, ip2Arr := isValidIp(ip2)
	if !b3 {
		fmt.Println(1)
		return
	}

	for i := 0; i < 4; i++ {
		if ip1Arr[i]&nsArr[i] != ip2Arr[i]&nsArr[i] {
			fmt.Println(2)
			return
		}
	}

	fmt.Println(0)
}

func isValidIp(ip string) (bool, []int) {
	var ipArr []int
	a := strings.Split(ip, ".")
	if len(a) != 4 {
		return false, ipArr
	}

	for i := 0; i < len(a); i++ {
		n, err := strconv.Atoi(a[i])
		if err != nil {
			return false, ipArr
		}

		if n < 0 || n > 255 {
			return false, ipArr
		}

		ipArr = append(ipArr, n)
	}

	return true, ipArr
}

func isValidNetmask(nsArr []int) bool {
	var sum uint32
	for i := 0; i < 4; i++ {
		n := nsArr[i] << ((3 - i) * 8)
		sum += uint32(n)
	}

	if sum == 0 || sum == 0xffffff {
		return false
	}

	return sum&(^sum+1) == ^sum+1
}
```

### 火车进站

### 公共子串计算
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getSubStr(s, substr string, index int) string {
	if strings.Contains(s, substr) {
		return substr
	}

	if index == 0 {
		return getSubStr(s, substr[1:], 1)
	}
	return getSubStr(s, substr[:len(substr)-1], 0)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var s1, s2 string
		s1 = scanner.Text()

		if scanner.Scan() {
			s2 = scanner.Text()
		}

		if len(s1) > len(s2) {
			s1, s2 = s2, s1
		}

		maxSubStr := getSubStr(s2, s1, 0)
		fmt.Println(len(maxSubStr))
	}
}
```

### 求解立方根
``` go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getCubValue(num int) float64 {
	var min, max float64
	min = 0
	max = float64(num)

	for max-min > 0.00000001 {
		mid := (max + min) / 2
		if mid*mid*mid > float64(num) {
			max = mid
		} else {
			min = mid
		}
	}

	return min
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		res := getCubValue(num)
		fmt.Printf("%.1f", res)
	}
}
```