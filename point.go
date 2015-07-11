package maze

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
