package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func CheckErrAndExit(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(0)
	}
}

func ParseArgs(width, height, sleeptime, percent *int) {
	for key, val := range os.Args {
		switch val {
		case "-h":
			if key+1 <= len(os.Args) {
				num, err := strconv.Atoi(os.Args[key+1])
				CheckErrAndExit(err)
				*height = num
			}
		case "-w":
			if key+1 <= len(os.Args) {
				num, err := strconv.Atoi(os.Args[key+1])
				CheckErrAndExit(err)
				*width = num
			}
		case "-s":
			if key+1 <= len(os.Args) {
				num, err := strconv.Atoi(os.Args[key+1])
				CheckErrAndExit(err)
				*sleeptime = num
			}
		case "-p":
			if key+1 <= len(os.Args) {
				num, err := strconv.Atoi(os.Args[key+1])
				CheckErrAndExit(err)
				*percent = num
			}
		}
		if *width == 0 {
			*width = 10
		}
		if *height == 0 {
			*height = 10
		}
		if *sleeptime == 0 {
			*sleeptime = 100
		}
		if *percent == 0 {
			*percent = 50
		}
	}
}

func InitCells(width, height, percent int) [][]Cell {
	cells := make([][]Cell, height)
	for key := range cells {
		cells[key] = make([]Cell, width)
		for pos, _ := range cells[key] {
			if rand.Intn(100) <= percent {
				cells[key][pos].Alive()
				cells[key][pos].Step()
			}
		}
	}
	return cells
}

func PrintImage(cells [][]Cell) {
	var TopAndBottomLine strings.Builder
	var Image strings.Builder
	for i := 0; i < len(cells[0])*2+2; i++ {
		TopAndBottomLine.WriteRune('-')
	}
	TopAndBottomLine.WriteRune('\n')
	Image.WriteString("\x1b[H\x1b[J")
	Image.WriteString(TopAndBottomLine.String())
	for _, line := range cells {
		Image.WriteRune('|')
		for _, cell := range line {
			if cell.IsAlive() {
				Image.WriteString("# ")
			} else {
				Image.WriteString("  ")
			}
		}
		Image.WriteString("|\n")
	}
	Image.WriteString(TopAndBottomLine.String())
	fmt.Println(Image.String())
}

func CheckCellsState(currentX, currentY int, cells [][]Cell) int {
	aliveCells := 0
	for i := currentY - 1; i <= currentY+1 && i < len(cells); i++ {
		if i < 0 {
			continue
		}
		for j := currentX - 1; j <= currentX+1 && j < len(cells[0]); j++ {
			if j < 0 {
				continue
			}
			if currentX == j && currentY == i {
				continue
			}
			if cells[i][j].IsAlive() {
				aliveCells++
			}
		}
	}
	return aliveCells
}

func ChangeCellsState(cells [][]Cell) {
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[0]); j++ {
			aliveCells := CheckCellsState(j, i, cells)
			if cells[i][j].IsAlive() {
				if aliveCells >= 2 && aliveCells <= 3 {
					cells[i][j].Alive()
				} else {
					cells[i][j].Kill()
				}
			} else {
				if aliveCells == 3 {
					cells[i][j].Alive()
				}
			}
		}
	}
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[0]); j++ {
			cells[i][j].Step()
		}
	}
}

func main() {
	var width, height, sleeptime, percent int
	ParseArgs(&width, &height, &sleeptime, &percent)
	cells := InitCells(width, height, percent)
	for {
		ChangeCellsState(cells)
		PrintImage(cells)
		time.Sleep(time.Millisecond * time.Duration(sleeptime))
	}
}
