package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

type Maze struct {
	Height int
	Width  int
	Board  []int
}

func process(dir string) {

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

	fmt.Printf(
		"Length: %d, Width: %d, squares: %d",
		maze.Height,
		maze.Width,
		maze.Height*maze.Width,
	)
}

func getMazeBoard(file io.Reader) (Maze, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return Maze{}, err
	}

	bounds := img.Bounds()

	width, height := bounds.Max.X, bounds.Max.Y

	fmt.Printf("width=%d, height=%d", width, height)

	var arr []int

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			r, g, b, a := img.At(x, y).RGBA()

			fmt.Printf("r=%d, g=%d, b=%d, a=%d \n", r/257, g/257, b/257, a/257)

			if r == 0 && g == 0 && b == 0 {
				// if pixel is black
				arr = append(arr, 1)
			} else {
				// if pixel is white
				arr = append(arr, 0)
			}
		}
	}

	return Maze{
		Height: height,
		Width:  width,
		Board:  arr,
	}, err
}
