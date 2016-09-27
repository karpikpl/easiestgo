package io

import (
	"bufio"
	"io"
	"strconv"
)

func ExecuteActionOnEachLine(r io.Reader, action func(int)) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		n, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}

		if n == 0 {
			// 0 means end of input
			return
		}

		action(n)
	}
}
