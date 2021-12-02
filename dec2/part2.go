package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	depth := 0
	progress := 0
	aim := 0
	for scanner.Scan() {
		command := scanner.Text()
		s := strings.Split(command, " ")
		action := s[0]
		val, _ := strconv.Atoi(s[1])
		switch action {
		case "forward":
			progress += val
			depth += aim * val
			break
		case "up":
			aim -= val
			break
		case "down":
			aim += val
			break
		}
	}
	fmt.Println(depth * progress)
}
