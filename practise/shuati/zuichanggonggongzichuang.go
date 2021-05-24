package main

import (
	"fmt"
)

func main() {
	for {
		var s1, s2 string
		n, _ := fmt.Scan(&s1, &s2)
		if n == 0 {
			break
		}

		if len(s1) > len(s2) {
			s1, s2 = s2, s1
		}

		// 		fmt.Println(getLongestSubStr(s1, s2))
		fmt.Println(getSubString(s1, s2))
	}
}

func getLongestSubStr(s1, s2 string) string {
	ls1, ls2 := len(s1), len(s2)
	max, maxIndex := -1, -1
	dp := make([][]int, ls1+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]int, ls2+1)
	}
	for i := 1; i < ls1+1; i++ {
		for j := 1; j < ls2+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > max {
					max = dp[i][j]
					maxIndex = i
				}
			}
		}
	}
	if max > 0 {
		return s1[maxIndex-max : maxIndex]
	}
	return ""
}

func getSubString(str1, str2 string) string {
	var maxSubStr string
	for begin1 := 0; begin1 < len(str1)-1; begin1++ {
		for begin2 := 0; begin2 < len(str2)-1; begin2++ {
			if str1[begin1] != str2[begin2] {
				continue
			}
		isEnd:
			for end1, end2 := begin1+1, begin2+1; end1 < len(str1) && end2 < len(str2); end1, end2 = end1+1, end2+1 {
				if str1[end1] != str2[end2] {
					if str1[end1-1] == str2[end2-1] && len(maxSubStr) < end1-begin1 {
						maxSubStr = str1[begin1:end1]
					}
					break isEnd
				}
			}
		}
	}
	return maxSubStr
}
