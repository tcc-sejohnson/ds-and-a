package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stringInput := getInput()
	highest, nextHighest := getMaxNums(stringInput)
	fmt.Print(highest * nextHighest)
}

func getInput() []string {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n') // We don't actually need the length
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(input), " ")
}

func getMaxNums(nums []string) (highest, nextHighest int64) {
	for _, numAsString := range nums {
		num, err := strconv.Atoi(numAsString)
		if err != nil {
			panic(err)
		}

		i64num := int64(num)
		if i64num > highest {
			nextHighest = highest
			highest = i64num
		} else if i64num > nextHighest {
			nextHighest = i64num
		}
	}
	return
}
