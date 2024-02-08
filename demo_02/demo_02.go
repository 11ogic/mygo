package main

import "fmt"

func printMulti(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v  ", j, i, i*j)
		}
		fmt.Print("\n")
	}
}

func printPyramid(layer int) {
	for i := 0; i < layer; i++ {
		// Print space first
		for k := 1; k <= layer-i; k++ {
			fmt.Printf(" ")
		}
		// Print start next
		for j := 1; j <= 1+i*2; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	var (
		multiParam int
		layerParam int
	)

	fmt.Println("Please enter your number")
	fmt.Scanln(&multiParam)
	printMulti(multiParam)

	fmt.Println("Please enter your layer")
	fmt.Scanln(&layerParam)
	printPyramid(layerParam)
}
