package main

import (
	"fmt"
	"github.com/karpikpl/easiestgo/pkg/io"
	"os"
)

func main() {
	io.ExecuteActionOnEachLine(os.Stdin, Solve)
}

func Solve(n int) {
	sum := SumDigits(n)
	var sumP int

	for i := 11; ; i++ {
		sumP = SumDigits(i * n)

		if sum == sumP {
			fmt.Println(i)
			return
		}
	}
}

func SumDigits(n int) int {
	sum := 0

	for n > 0 {
		sum += n - ((int)(n/10) * 10)
		n /= 10
	}

	return sum
}
