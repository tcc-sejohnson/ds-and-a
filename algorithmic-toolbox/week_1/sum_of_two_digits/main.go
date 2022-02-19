package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	numsSplit := strings.Split(nums, " ")
	if len(numsSplit) != 2 {
		panic("Did not receive two numbers on the command line.")
	}

	num1, err := strconv.Atoi(numsSplit[0])
	if err != nil {
		panic(err)
	}

	num2, err := strconv.Atoi(strings.TrimSpace(numsSplit[1]))
	if err != nil {
		panic(err)
	}

	fmt.Print(num1 + num2)
}
