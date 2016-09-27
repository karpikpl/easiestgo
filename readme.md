# About
task from https://open.kattis.com/problems/easiest

## Input
The input consists of several test cases. Each case is described by a single line containing one positive integer number `N,1≤N≤100000N` The last test case is followed by a line containing zero.

### How to Test
http://blog.stretchr.com/2014/03/05/test-driven-development-specifically-in-golang/

```
$ go test ./... -v
$ go test ./easiest -test.coverprofile=coverageReport.out -test.v=true
```

### Benchmarking

```
$ go test ./... -test.bench=".*" -test.v=true
```

### How to build

```
$ go build ./...
```

### How to format

```
$ go fmt ./...
```

### How to run

```
$ go run main\main.go < sample.in > answer.ans
```

### Sources
https://github.com/karpikpl/easiestgo.git
