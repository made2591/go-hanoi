package main

import (
	"fmt"
	m "github.com/skelterjohn/go.matrix"
)

func cell(d int, m int) string {

	r := ""

	s := int(m/2) - int(d/2)

	for i := 0; i < m; i++ {

		if i < s || d == 0 {
			r += " "
		} else {
			r += "*"
			d--
		}

	}

	return r + " |"

}

func move(h *m.DenseMatrix, s int, d int) {

	v := 0

	for r := 0; r < h.Rows(); r++ {

		v = int(h.Get(r, s))
		if v > 0 {

			h.Set(r, s, 0.0)
			break

		}

	}

	for r := h.Rows() - 1; r >= 0; r-- {

		if h.Get(r, d) == 0 {

			h.Set(r, d, float64(v))
			break

		}

	}

}

func resolve(h *m.DenseMatrix, s int, d int, a int, m int, n int, p int) int {

	r := p

	if n == 1 {

		move(h, s, d)
		fmt.Println(stringfy(h, m))
		return r + 1

	} else {

		r += resolve(h, s, a, d, m, n-1, p)

		move(h, s, d)
		fmt.Println(stringfy(h, m))

		return r + resolve(h, a, d, s, m, n-1, p)

	}

}

func prepare(nd int, nc int, cs int) *m.DenseMatrix {

	dd := []float64{}

	for r := 0; r < nd; r++ {
		for c := 0; c < nc; c++ {
			if c == cs {
				dd = append(dd, float64(r+1))
			} else {
				dd = append(dd, 0)
			}
		}
	}

	return m.MakeDenseMatrix(dd, nd, nc)

}

func hanoi(nd int, nc int, cs int, cd int) int {

	h := prepare(nd, nc, cs)

	ca := 0
	for c := 0; c < h.Cols(); c++ {
		if c != cs && c != cd {
			ca = c
			break
		}
	}
	fmt.Println(stringfy(h, nd*2))

	return resolve(h, cs, cd, ca, nd*2, nd, 0)

}

func stringfy(h *m.DenseMatrix, m int) string {

	hs := "|"
	for c := 0; c < h.Cols()*m+(h.Cols()*3)-1; c++ {
		hs += "-"
	}
	hs += "|\n"

	for r := 0; r < h.Rows(); r++ {
		hs += "| "
		for c := 0; c < h.Cols(); c++ {
			hs += cell(int(h.Get(r, c))*2, m) + " "
		}
		hs += " \n"
	}

	hs += "|"
	for c := 0; c < h.Cols()*m+(h.Cols()*3)-1; c++ {
		hs += "-"
	}
	hs += "|\n"

	return hs

}

func main() {

	s := []int{}

	s = append(s, hanoi(3, 3, 0, 1)*2)
	s = append(s, hanoi(3, 4, 0, 1)*2)
	s = append(s, hanoi(3, 5, 0, 1)*2)
	s = append(s, hanoi(4, 3, 0, 1)*2)
	s = append(s, hanoi(5, 3, 0, 1)*2)
	s = append(s, hanoi(7, 3, 0, 1)*2)

	fmt.Println(s)

}
