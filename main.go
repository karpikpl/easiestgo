package main

import (
	"easiest/io"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			fmt.Println(p)
			return
		}
	}
}

func SumDigits(n int) int {
	nString := strconv.Itoa(n)
	digits := strings.SplitN(nString, "", -1)

	sum := 0
	for _, digit := range digits {
		i, err := strconv.Atoi(digit)

		if err != nil {
			panic(nil)
		}
		sum += i
	}

	return sum
}
