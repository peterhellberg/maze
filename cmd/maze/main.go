package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/peterhellberg/maze"
)

var (
	width  = flag.Int("width", 17, "The width of the maze")
	height = flag.Int("height", 8, "The height of the maze")
	seed   = flag.Int64("seed", 3, "The starting seed for the maze")
)

func main() {
	// Parse the command line flags
	flag.Parse()

	// Seed the random number generator
	rand.Seed(*seed)

	// Create a new maze
	m := maze.New(*width, *height)

	// Print the maze
	fmt.Println(m)
}
