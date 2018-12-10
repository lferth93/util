package torus

import (
	"testing"
)

func TestType(t *testing.T) {
	torus := New(0, 5, 5)
	if torus.Rows() == 5 && torus.Cols() == 5 {
		for _, r := range torus.data {
			for _, e := range r {
				if _, ok := e.(int); !ok {
					t.Error("diferent types")
				}
			}
		}
	} else {
		t.Error("Diferent size")
	}
}

func TestTypeError(t *testing.T) {
	torus := New(false, 5, 5)
	if err := torus.Init(3); err == nil {
		t.Error("Diferent types not detected")
	}
	if err := torus.Init(true); err != nil {
		t.Error("Diferent types detected")
	}
}
