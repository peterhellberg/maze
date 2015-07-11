package maze

import "math/rand"

var (
	// StartRune is used for the starting point in the maze
	StartRune = 'S'

	// FinishRune is used for the finish point in the maze
	FinishRune = 'F'

	// WallRune is used for all walls in the maze
	WallRune = '*'

	// EmptyRune is used for all space in the maze
	EmptyRune = ' '
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
			maze[row][ch] = WallRune
		}
	}

	p := &point{x: rand.Intn(w - 2), y: rand.Intn(h - 2)}
	maze[p.x][p.y] = StartRune

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

		if inMaze(opp.x, opp.y, w-2, h-2) && maze[opp.x][opp.y] == WallRune {
			maze[wall.x][wall.y] = EmptyRune
			maze[opp.x][opp.y] = EmptyRune
			walls = append(walls, adjacents(opp, maze)...)
			f = opp
		}
	}

	maze[f.x][f.y] = FinishRune
	bordered := make([][]rune, len(maze)+2)

	for r := range bordered {
		bordered[r] = make([]rune, len(maze[0]))

		for c := range bordered[r] {
			if r == 0 || r == len(maze)+1 || c == 0 || c == len(maze[0])+1 {
				bordered[r][c] = WallRune
			} else {
				bordered[r][c] = maze[r-1][c-1]
			}
		}
	}

	return bordered
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

func (m Maze) Width() int {
	return len(m)
}

func (m Maze) Height() int {
	return len(m[0])
}

func adjacents(p *point, m Maze) []*point {
	var res []*point

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if (i == 0 && j == 0) || (i != 0 && j != 0) {
				continue
			}

			if !inMaze(p.x+i, p.y+j, m.Width(), m.Height()) {
				continue
			}

			if m[p.x+i][p.y+j] == WallRune {
				res = append(res, &point{p.x + i, p.y + j, p})
			}
		}
	}

	return res
}

func inMaze(x, y int, w, h int) bool {
	return x >= 0 && x < w && y >= 0 && y < h
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
