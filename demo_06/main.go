package main

import "fmt"

func main() {
	arr := [2]string{"你", "好"}
	fmt.Printf("arr point: %p, value: %v \n", &arr, arr)
	fmt.Printf("arr[0] point: %p, value: %v \n", &arr[0], arr[0])
	fmt.Printf("arr[1] point: %p, value: %v \n", &arr[1], arr[1])

}
????