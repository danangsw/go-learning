package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (r [][]uint8) {
	return picSomething(allocateSlice(dx, dy))
}

func allocateSlice(dx, dy int) (r [][]uint8) {
	// alokasikan array 2 dimensi
	r = make([][]uint8, dy)
	
	for i:=0; i<dy; i++ {
		r[i] = make([]uint8, dx)
	}
	
	return r
}

func picSomething(r [][]uint8) [][]uint8 {
	// bentuk gambar
	for i := range r {
        for j := range r[i] {
            r[i][j] = uint8(((j+i)/2)^(i+j))
        }
    }
	
	return r
}

func main() {
	pic.Show(Pic)
}
