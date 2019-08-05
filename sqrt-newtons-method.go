package main

import (
	"fmt"
)

func Abs(x float64) float64 {
	if x > 0 {
		return x
	} else {
		return x * -1
	}
}

func Sqrt(x float64) float64 {
	z := 1.0					//The initial value
	tolerance := 0.00000000001	//The digit accuracy is desired 
	
	//maxIteration := 20		//Don't allow the iterations to continue indefinitely
	for true {
		z0 := z
		
		z -= (z * z - x) / (2*z)
		
		//fmt.Println(z0, z, z0 - z)
		
		if Abs(z - z0) < tolerance {
			break
		}
	}
	
	return z;
}

func main() {
	var res = Sqrt(25)
	fmt.Println(res, res*res)
}
