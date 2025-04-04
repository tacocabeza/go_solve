package main

import (
	"fmt"
	"os"
)

func main() {

	dir := os.Args[1]

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println(dir, "does not exist!")
		os.Exit(1)
	}

	fmt.Println(dir)

	maze, err := process(dir)

	if err != nil {
		fmt.Println("Could not process Maze")
	}

	fmt.Printf("(main.go) - maze board:%d\n", maze.Board)

	solve(maze)

}
