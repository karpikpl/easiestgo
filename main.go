package main

import (
	"easiest/io"
	"fmt"
	"os"
)

func main() {
	io.ExecuteActionOnEachLine(os.Stdin, Solve)
}

func Solve(n int) {
	if n == 0 {
		return
	}

	sum := SumDigits(n)

	for i := 11; ; i++ {
		p := i * n
		sumP := SumDigits(p)

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
