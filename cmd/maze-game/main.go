package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/maze"
)

var (
	width  = flag.Int("w", 24, "Width of the screen")
	height = flag.Int("h", 12, "Height of the screen")
	scale  = flag.Int("s", 36, "Scaling factor")

	Orange = color.RGBA{0xFF, 0x66, 0x00, 0xFF}
	Blue2  = color.RGBA{0x14, 0x5d, 0x7b, 0xff}
	Blue6  = color.RGBA{0x06, 0x1f, 0x29, 0xff}
	Gray   = color.RGBA{0xCF, 0xCF, 0xCF, 0xFF}
)

func main() {
	flag.Parse()

	if *width < 10 || *height < 10 || *scale < 10 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	game := newGame(*width, *height)

	if err := game.Run(*width, *height, *scale, "Maze"); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	Maze             maze.Maze
	Image            *image.RGBA
	Bounds           image.Rectangle
	Frame            *ebiten.Image
	DrawImageOptions *ebiten.DrawImageOptions
	Score            int
	W                int
	H                int
	X                int
	Y                int
}

func newGame(w, h int) *Game {
	g := &Game{W: w, H: h, DrawImageOptions: &ebiten.DrawImageOptions{}}
	g.newMaze()

	return g
}

func (g *Game) newMaze() {
	g.Maze = maze.New(g.W, g.H)
	g.Image = image.NewRGBA(image.Rect(0, 0, g.W, g.H))
	g.Bounds = g.Image.Bounds()

	for x := 0; x < g.W; x++ {
		for y := 0; y < g.H; y++ {
			switch g.Maze[x][y] {
			case maze.Start:
				g.X = x
				g.Y = y
				g.Image.Set(x, y, Gray)
			case maze.Finish:
				g.Image.Set(x, y, Blue2)
			case maze.Wall:
				g.Image.Set(x, y, Blue6)
			case maze.Empty:
				g.Image.Set(x, y, Gray)
			}
		}
	}
}

func (g *Game) Run(width, height, scale int, title string) error {
	ticker := time.NewTicker(time.Millisecond * 70)
	go func() {
		for range ticker.C {
			if ebiten.IsKeyPressed(ebiten.KeyR) {
				g.Score = 0
				g.newMaze()
			}

			if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.canGoLeft() {
				g.X--
			}

			if ebiten.IsKeyPressed(ebiten.KeyRight) && g.canGoRight() {
				g.X++
			}

			if ebiten.IsKeyPressed(ebiten.KeyUp) && g.canGoUp() {
				g.Y--
			}

			if ebiten.IsKeyPressed(ebiten.KeyDown) && g.canGoDown() {
				g.Y++
			}
		}
	}()
	defer ticker.Stop()

	return ebiten.Run(g.update, width, height, scale, title)
}

func (g *Game) updateFrame() {
	m := image.NewRGBA(g.Bounds)
	draw.Draw(m, g.Bounds, g.Image, g.Bounds.Min, draw.Src)

	m.Set(g.X, g.Y, Orange)

	if f, err := ebiten.NewImageFromImage(m, ebiten.FilterLinear); err == nil {
		g.Frame = f
	}
}

func (g *Game) update(screen *ebiten.Image) error {
	if g.Frame == nil ||
		ebiten.IsKeyPressed(ebiten.KeyUp) ||
		ebiten.IsKeyPressed(ebiten.KeyDown) ||
		ebiten.IsKeyPressed(ebiten.KeyLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.updateFrame()
	}

	if err := screen.DrawImage(g.Frame, g.DrawImageOptions); err != nil {
		return err
	}

	if g.Maze[g.X][g.Y] == maze.Finish {
		g.Score += (g.W * g.H) / 100
		g.newMaze()
		fmt.Println("Started new maze, current score:", g.Score)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return errors.New("exited")
	}

	return nil
}

func (g *Game) canGoLeft() bool {
	if g.X < 1 {
		return false
	}

	if t := g.Maze[g.X-1][g.Y]; t != maze.Wall {
		return true
	}

	return false
}

func (g *Game) canGoRight() bool {
	if g.X+1 == g.W {
		return false
	}

	if t := g.Maze[g.X+1][g.Y]; t != maze.Wall {
		return true
	}

	return false
}

func (g *Game) canGoUp() bool {
	if g.Y < 1 {
		return false
	}

	if t := g.Maze[g.X][g.Y-1]; t != maze.Wall {
		return true
	}

	return false
}

func (g *Game) canGoDown() bool {
	if g.Y+1 == g.H {
		return false
	}

	if t := g.Maze[g.X][g.Y+1]; t != maze.Wall {
		return true
	}

	return false
}
