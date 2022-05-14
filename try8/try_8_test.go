package main_test

import "testing"

var cases = []struct {
	name string
	in   int
	out  bool
}{
	{name: "success odd", in: 1, out: false},
	{name: "success even", in: 2, out: true},
}

func TestIsEven(t *testing.T) {

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := main.IsEven(c.in); c.out != actual {
				t.Errorf("なんか違う")
			}
		})

	}
}
