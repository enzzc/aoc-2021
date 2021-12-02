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
	for scanner.Scan() {
		command := scanner.Text()
		s := strings.Split(command, " ")
		action := s[0]
		val, _ := strconv.Atoi(s[1])
		switch action {
		case "forward":
			progress += val
			break
		case "up":
			depth -= val
			break
		case "down":
			depth += val
			break
		}
	}
	fmt.Println(depth * progress)
}
