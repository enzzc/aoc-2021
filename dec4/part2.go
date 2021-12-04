package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const gridWidth = 5

type Grid [gridWidth * gridWidth]int

func main() {
	/* Step 1: populate grids */
	var grids []Grid
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	var outcomes []int
	line := strings.Split(scanner.Text(), ",")
	for _, part := range line {
		n, _ := strconv.Atoi(part)
		outcomes = append(outcomes, n)
	}

	gridIdx := 0
	var grid Grid
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		nums := strings.Fields(line)
		for _, part := range nums {
			n, _ := strconv.Atoi(part)
			grid[gridIdx] = n
			gridIdx++
		}
		if gridIdx >= gridWidth*gridWidth {
			gridIdx = 0
			grids = append(grids, grid)
		}
	}

	/* Step 2: play bingo */
	won := make(map[int]bool)
	t0 := time.Now()
	lastGrid := -1
	lastCall := -1
	steps := 0
	for _, outcome := range outcomes {
		for i, _ := range grids {
			_, present := won[i]
			if present {
				continue
			}
			steps++
			checkNumbersOnGrid(&grids[i], outcome)
			if isBingo(&grids[i]) {
				won[i] = true
				lastGrid = i
				lastCall = outcome
			}
		}
	}
	fmt.Println("Steps:", steps, "Took:", time.Since(t0))
	accUnchecked := 0
	for _, n := range grids[lastGrid] {
		if n > 0 {
			accUnchecked += n
		}
	}
	fmt.Println(accUnchecked * lastCall)
}

// We check a number by doing n = -n - 1
func checkNumbersOnGrid(grid *Grid, outcome int) {
	for i, n := range grid {
		if n == outcome {
			grid[i] = -grid[i] - 1
		}
	}
}

func isBingo(grid *Grid) bool {
	/* Row-wise traversals */
	for row := 0; row < gridWidth; row++ { // For each row in the matrix
		checked := 0
		for k := row * gridWidth; k < gridWidth*row+gridWidth; k++ { // For each item in row
			if grid[k] >= 0 {
				break
			}
			checked++
			if checked == gridWidth {
				return true
			}
		}
	}
	/* Column-wise traversals */
	for col := 0; col < gridWidth; col++ { // For each col in the matrix
		checked := 0
		for k := col; k < gridWidth*gridWidth+col; k += gridWidth {
			if grid[k] >= 0 {
				break
			}
			checked++
			if checked == gridWidth {
				return true
			}
		}
	}
	return false
}
