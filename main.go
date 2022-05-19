package main

func AddTwo(x float64) (result float64) {
	return Add(x, 2)
}

func Add(x float64, y float64) (result float64) {
	return x + y
}

// How to see test coverage
// First, run command below to make a coverage file
// go test -coverprofile=coverage.out
// then, run command below to see the result of coverage file
// go tool cover -html coverage.out

func main() {}
