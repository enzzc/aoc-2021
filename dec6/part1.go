package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	timerStates = 9
	targetDay   = 80
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, ",")
	var initialState []int
	for _, num := range parts {
		n, _ := strconv.Atoi(num)
		initialState = append(initialState, n)
	}

	population := 0
	var counters [timerStates]int
	for _, timer := range initialState {
		counters[timer] += 1
		population++
	}

	t0 := time.Now()
	for i := 1; i <= targetDay; i++ {
		var newCounters [timerStates]int
		for timer, count := range counters {
			if timer == 0 {
				population += count
				newCounters[8] += count
				newCounters[6] += count
			} else {
				newCounters[timer-1] += count
			}
		}
		counters = newCounters
	}

	fmt.Println(population, "Took:", time.Since(t0))
}
