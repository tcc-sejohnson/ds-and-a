package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Goal: Find the least common multiple of two integers.
// LCM = a*b / GCD(a, b)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	rawNums := strings.Split(input, " ")
	nums := make([]int, 2)
	for i, rawNum := range rawNums {
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	fmt.Print(LCM(nums[0], nums[1]))
}

func LCM(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
