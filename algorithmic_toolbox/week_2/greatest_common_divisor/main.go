package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Goal: Find the greatest common denominator of two integers.

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
	fmt.Print(GCD(nums[0], nums[1]))
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
