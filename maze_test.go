package maze

import "testing"

func TestInMaze(t *testing.T) {
	for i, tt := range []struct {
		x, y int
		w, h int
		b    bool
	}{
		{10, 10, 10, 10, false},
		{20, 30, 10, 20, false},
		{10, 10, 50, 50, true},
		{10, 10, 11, 50, true},
	} {
		if got := inMaze(tt.x, tt.y, tt.w, tt.h); got != tt.b {
			t.Fatalf(`T%d: inMaze(%d, %d, %d, %d) = %v, want %v`,
				i, tt.x, tt.y, tt.w, tt.h, got, tt.b)
		}
	}
}

func TestPointOpposite(t *testing.T) {
	for i, tt := range []struct {
		x1, y1 int
		x2, y2 int
		x3, y3 int
	}{
		{10, 20, 12, 15, 8, 20},
		{100, 15, 100, 9, 100, 21},
		{15, 120, 10, 230, 20, 120},
	} {
		p := &point{x: tt.x1, y: tt.y1, p: &point{x: tt.x2, y: tt.y2}}

		o := p.opposite()
		if o == nil {
			t.Fatal(`got unexpected nil`)
		}

		if p != o.p {
			t.Fatal(`p != o.p`)
		}

		if o.x != tt.x3 {
			t.Fatalf(`T%d: o.x = %d, want %d`, i, o.x, tt.x3)
		}

		if o.y != tt.y3 {
			t.Fatalf(`T%d: o.y = %d, want %d`, i, o.y, tt.y3)
		}
	}
}
