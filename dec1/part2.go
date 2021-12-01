package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	last, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	beforeLast, _ := strconv.Atoi(scanner.Text())
	lastSum := 0
	counter := 0
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		sum := n + last + beforeLast
		if lastSum != 0 && sum > lastSum {
			counter++
		}
		last = beforeLast
		beforeLast = n
		lastSum = sum
	}
	fmt.Println(counter)
}
