package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//
//	Submarine struct
//

type Submarine struct {
	horizontalPos int
	depth int
}

func (sub *Submarine) Forward(x int) {
	sub.horizontalPos += x
}

func (sub *Submarine) Down(x int) {
	sub.depth += x
}

func (sub *Submarine) Up(x int) {
	sub.depth -= x
}

func (sub *Submarine) Multiply() int {
	return sub.depth * sub.horizontalPos
}

//
//	Main
//

func ExecuteLine(sub *Submarine, line string) error {
	split := strings.Split(line, " ")
	order := split[0]
	amount, err := strconv.Atoi(split[1])
	if err != nil {
		return errors.New("Failed to convert to int: " + err.Error())
	}

	switch order {
	case "forward":
		sub.Forward(amount)
	case "down":
		sub.Down(amount)
	case "up":
		sub.Up(amount)
	default:
		log.Println("Invalid order, ignoring.")
	}

	return nil
}

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Not enough argument given. Specify an input file.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Failed to open file: " + err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var sub Submarine
	for scanner.Scan() {
		err := ExecuteLine(&sub, scanner.Text())
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	fmt.Println("Forwards * Depth = ", sub.Multiply())
}
