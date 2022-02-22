package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Goal: Calculate the n-th fibonacci number's last digit in a reasonable amount of time.

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	fmt.Print(Fib(n))
}

func Fib(n int) int {
	memo := make(map[int]int)
	return FibRecurse(n, memo)
}

func FibRecurse(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if fibNum, exists := memo[n]; exists {
		return fibNum
	}
	fibNum := (FibRecurse(n-1, memo) + FibRecurse(n-2, memo)) % 10
	memo[n] = fibNum
	return fibNum
}
