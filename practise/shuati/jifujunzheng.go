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
	for {
		var n int
		if scanner.Scan() {
			n, _ = strconv.Atoi(scanner.Text())
		} else {
			return
		}

		negCount, posCount := 0, 0
		sum := 0
		var v int
		for i := 0; i < n; i++ {
			if scanner.Scan() {
				v, _ = strconv.Atoi(scanner.Text())
			}
			if v < 0 {
				negCount++
				continue
			}
			if v > 0 {
				posCount++
				sum += v
			}
		}

		var avg float64
		if posCount > 0 {
			avg = float64(sum) / float64(posCount)
		}
		fmt.Printf("%d %.1f\n", negCount, avg)
	}
}
