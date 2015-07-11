package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	for i, tt := range []struct {
		width  string
		height string
		seed   string
		out    string
	}{
		{"5", "5", "1", expectedMazes["5x5:1"]},
		{"10", "5", "2", expectedMazes["10x5:2"]},
	} {
		out, err := execGo("run", "main.go", "-seed", tt.seed, "-width", tt.width, "-height", tt.height)
		if err != nil {
			t.Fatalf(`T%d: unexpected error: %v`, i, err)
		}

		if out != tt.out+"\n" {
			t.Fatalf("T%d: unexpected maze:\n%s\nwanted:\n\n%s", i, out, tt.out)
		}
	}
}

var (
	expectedMazes = map[string]string{
		"5x5:1": strings.Join([]string{
			"*****",
			"*F S*",
			"*** *",
			"*   *",
			"*****",
		}, "\n"),
		"10x5:2": strings.Join([]string{
			"**********",
			"*  S    **",
			"* * *** **",
			"* *   *F**",
			"**********",
		}, "\n"),
	}
)

func execGo(args ...string) (string, error) {
	out, err := exec.Command("go", args...).CombinedOutput()

	return string(out), err
}
