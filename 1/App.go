package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Not enough arguments provided.")
		return
	}

	file := os.Args[1]
	contentBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		return
	}
	content := string(contentBytes)
	lines := strings.Split(content, "\n")

	var numbers []int
	for _, line := range lines {
		line = strings.Trim(line, " \n\r\t")
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		numbers = append(numbers, num)
	}

	increases := 0
	decreases := 0

	for i := 3; i < len(numbers); i++ {
		var one, two, three = numbers[i], numbers[i - 1], numbers[i - 2]
		var lastOne, lastTwo, lastThree = numbers[i - 1], numbers[i - 2], numbers[i - 3]
		sum := one + two + three
		lastSum := lastOne + lastTwo + lastThree

		fmt.Println("CURRENT One = " + strconv.Itoa(one),
			"Two = " + strconv.Itoa(two),
			"Three = " + strconv.Itoa(three))
		fmt.Println("LAST One = " + strconv.Itoa(lastOne),
			"Two = " + strconv.Itoa(lastTwo),
			"Three = " + strconv.Itoa(lastThree) + "\n\n")

		if lastSum < sum {
			increases += 1
		} else if lastSum > sum {
			decreases += 1
		}
	}

	fmt.Println("Increases: " + strconv.Itoa(increases), "Decreases " + strconv.Itoa(decreases))
}
