package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Solution(328))
	fmt.Println(Solution(529))
	fmt.Println(Solution(9))
	fmt.Println(Solution(66561))
	fmt.Println(Solution(805306373))
	fmt.Println(Solution(74901729))

}

func Solution(N int) int {
	if N > 0 && N < int(math.Pow(2.0, 32.0)) {
		bin := strconv.FormatInt(int64(N), 2)
		first := strings.Index(bin, "1")
		next := first
		largest := 0
		for next >= 0 {
			if len(bin) >= next+1 {
				bin = bin[next+1:]
				next = strings.Index(bin, "1")
				if next >= 0 {
					if next-first > largest {
						largest = next - first
					}
					first = 0
				} else {
					break
				}
			}
		}
		return largest
	}

	return 0
}
