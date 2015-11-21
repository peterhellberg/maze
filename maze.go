package maze

import "math/rand"

var (
	// Start is used for the starting point in the maze
	Start = 'S'

	// Finish is used for the finish point in the maze
	Finish = 'F'

	// Wall is used for all walls in the maze
	Wall = '*'

	// Empty is used for all space in the maze
	Empty = ' '
)

// Maze holds the runes of a generated maze
type Maze [][]rune

// New generates a Maze using Prim's Algorithm
// https://en.wikipedia.org/wiki/Maze_generation_algorithm#Randomized_Prim.27s_algorithm
func New(w, h int) Maze {
	maze := make([][]rune, w-2)

	for row := range maze {
		maze[row] = make([]rune, h)
		for ch := range maze[row] {
			maze[row][ch] = Wall
		}
	}

	p := &point{x: rand.Intn(w - 2), y: rand.Intn(h - 2)}
	maze[p.x][p.y] = Start

	var f *point

	walls := adjacents(p, maze)

	for len(walls) > 0 {
		wall := walls[rand.Intn(len(walls))]

		for i, w := range walls {
			if w.x == wall.x && w.y == wall.y {
				walls = append(walls[:i], walls[i+1:]...)
				break
			}
		}

		opp := wall.opposite()

		if inMaze(opp.x, opp.y, w-2, h-2) && maze[opp.x][opp.y] == Wall {
			maze[wall.x][wall.y] = Empty
			maze[opp.x][opp.y] = Empty
			walls = append(walls, adjacents(opp, maze)...)
			f = opp
		}
	}

	maze[f.x][f.y] = Finish

	return borderedMaze(maze)
}

func borderedMaze(maze Maze) Maze {
	b := make([][]rune, len(maze)+2)

	for r := range b {
		b[r] = make([]rune, len(maze[0]))

		for c := range b[r] {
			if r == 0 || r == len(maze)+1 || c == 0 || c == len(maze[0])+1 {
				b[r][c] = Wall
			} else {
				b[r][c] = maze[r-1][c-1]
			}
		}
	}

	return b
}

func inMaze(x, y int, w, h int) bool {
	return x >= 0 && x < w && y >= 0 && y < h
}

func (m Maze) String() string {
	s := ""

	for y := 0; y < len(m[0]); y++ {
		for x := 0; x < len(m); x++ {
			s += string(m[x][y])
		}
		s += "\n"
	}

	return s
}

// Width returns the width of the maze
func (m Maze) Width() int {
	return len(m)
}

// Height returns the height of the maze
func (m Maze) Height() int {
	return len(m[0])
}

func (m Maze) include(x, y int) bool {
	return x >= 0 && x < m.Width() && y >= 0 && y < m.Height()
}

func adjacents(p *point, m Maze) []*point {
	var res []*point

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if (i == 0 && j == 0) || (i != 0 && j != 0) {
				continue
			}

			if !m.include(p.x+i, p.y+j) {
				continue
			}

			if m[p.x+i][p.y+j] == Wall {
				res = append(res, &point{p.x + i, p.y + j, p})
			}
		}
	}

	return res
}

type point struct {
	x int
	y int
	p *point
}

func (p *point) opposite() *point {
	if p.x != p.p.x {
		return &point{x: p.x + (p.x - p.p.x), y: p.y, p: p}
	}

	if p.y != p.p.y {
		return &point{x: p.x, y: p.y + (p.y - p.p.y), p: p}
	}

	return nil
}
