package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	dir := os.Args[1]

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println(dir, "does not exist!")
		os.Exit(1)
	}

	fmt.Println(dir)

	start := time.Now()
	maze, err := process(dir)
	elasped := time.Since(start)
	fmt.Printf("(main.go) - processing maze took %s\n", elasped)

	if err != nil {
		fmt.Println("Could not process Maze")
	}

	fmt.Printf(
		"(main.go) - start/end (%d,%d)/(%d/%d)",
		maze.Start.X,
		maze.Start.Y,
		maze.End.X,
		maze.End.Y,
	)

	start = time.Now()
	solve(maze)
	elasped = time.Since(start)
	fmt.Printf("(main.go) - solving maze took %s\n", elasped)

}
