package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

type Point struct {
	X int
	Y int
}

type Maze struct {
	Height      int
	Width       int
	Board       [][]int
	wasHere     [][]bool
	correctPath [][]bool
	Start       Point
	End         Point
}

func process(dir string) (Maze, error) {

	fmt.Println("processing file", dir)

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(dir)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	maze, err := getMazeBoard(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	fmt.Printf("len=%d cap=%d %v\n", len(maze.Board), cap(maze.Board), maze.Board)

	return maze, err
}

func getMazeBoard(file io.Reader) (Maze, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return Maze{}, err
	}

	bounds := img.Bounds()

	width, height := bounds.Max.X, bounds.Max.Y

	fmt.Printf("width=%d, height=%d", width, height)

	visited := make([][]bool, width)
	correctPath := make([][]bool, width)

	for i := range visited {
		visited[i] = make([]bool, height)
	}

	for i := range correctPath {
		correctPath[i] = make([]bool, height)
	}

	fmt.Printf("(process.go): visited? %t\n", visited)
	fmt.Printf("(process.go): correctPath? %t\n", correctPath)

	board := make([][]int, width)
	for i := range board {
		board[i] = make([]int, height)
	}

	startX := 0
	startY := 0
	foundStart := false

	endX := 0
	endY := 0

	for y := range board {
		for x := range board[y] {
			r, g, b, _ := img.At(x, y).RGBA()

			// fmt.Printf("r=%d, g=%d, b=%d, a=%d \n", r/257, g/257, b/257, a/257)

			if r/257 == 0 && g/257 == 0 && b/257 == 0 {
				// pixel is black
				board[x][y] = 1
			} else {
				// if pixel is white
				board[x][y] = 0
			}

			if y == 0 || x == width-1 || x == 0 || y == height-1 {
				if board[x][y] == 0 && foundStart {
					endX = x
					endY = y
					fmt.Printf("END: (%d, %d)\n", x, y)
				} else if board[x][y] == 0 && !foundStart {
					foundStart = true
					startX = x
					startY = y
					fmt.Printf("START: (%d, %d)\n", x, y)
				}
			}

		}
	}

	return Maze{
		Height:      height,
		Width:       width,
		Board:       board,
		wasHere:     visited,
		correctPath: correctPath,
		Start:       Point{X: startX, Y: startY},
		End:         Point{X: endX, Y: endY},
	}, err
}

func recursiveSolve(mazeToSolve Maze, x int, y int) bool {

	fmt.Printf("position: (%d, %d)\n", x, y)

	if x == mazeToSolve.End.X && y == mazeToSolve.End.Y {
		fmt.Printf("recursiveSolve: got to the end (%d, %d)\n", x, y)
		return true
	}

	if mazeToSolve.Board[x][y] == 1 || mazeToSolve.wasHere[x][y] == true {
		return false
	}

	mazeToSolve.wasHere[x][y] = true

	if x != 0 {
		if recursiveSolve(mazeToSolve, x-1, y) {
			mazeToSolve.correctPath[x][y] = true
			return true
		}
	}

	if x != mazeToSolve.Width-1 {
		if recursiveSolve(mazeToSolve, x+1, y) {
			mazeToSolve.correctPath[x][y] = true
			return true
		}
	}

	if y != 0 {
		if recursiveSolve(mazeToSolve, x, y-1) {
			mazeToSolve.correctPath[x][y] = true
			return true
		}
	}

	if y != mazeToSolve.Height-1 {
		if recursiveSolve(mazeToSolve, x, y+1) {
			mazeToSolve.correctPath[x][y] = true
			return true
		}
	}

	return false
}

func solve(mazeToSolve Maze) Maze {

	recursiveSolve(mazeToSolve, mazeToSolve.Start.X, mazeToSolve.Start.Y)

	return mazeToSolve
}
