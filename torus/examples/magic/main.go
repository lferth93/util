package main

import (
	"fmt"

	"github.com/lferth93/util/torus"
)

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	t := torus.New(0, n, n)
	x, y := t.Cols()/2, t.Rows()-1
	for i := 1; i <= n*n; i++ {
		t.Set(y, x, i)
		if t.Get(y+1, x+1).(int) > 0 {
			y--
		} else {
			x++
			y++
		}
	}
	fmt.Println(t)
}
