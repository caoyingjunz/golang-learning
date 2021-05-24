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
	r := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := r.ReadLine()
		if len(data) == 0 {
			break
		}
		str1 := strings.Fields(string(data))
		if len(str1) < 2 {
			break
		}
		n, _ := strconv.Atoi(str1[0])
		k, _ := strconv.Atoi(str1[1])

		imput2, _, _ := r.ReadLine()
		str2 := strings.Fields(string(imput2))
		arr := make([]int, 0, n)
		for i := 0; i < n; i++ {
			s, _ := strconv.Atoi(str2[i])
			arr = append(arr, s)
		}
		sort.Ints(arr)

		for i := 0; i < k; i++ {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Println()
	}

}
