package main

import "fmt"

func main() {
	loop(3)
	
	fmt.Println(c())
}

func c() (i int) {
	x := 0
    defer func() { x++ }()
    return 2
}

func loop(x int) {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		
		if(i == x) {
			return
		}
	}
	
	defer fmt.Println("done") 
}
