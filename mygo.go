package main

import "fmt"

func add(n *int) {
	*n++      // 改变地址的值
	p1 := n   // p1 赋值 参数n指针
	p2 := &p1 // p2 赋值 p1的指针
	p3 := &p2 // p3 赋值 p2的指针

	*p3 = p2 // p3 地址的值赋值 p2 地址(&p1) => n(地址)

	fmt.Printf("add num: %v \n", n)
	fmt.Printf("add num: %v \n", p1)
	fmt.Printf("add num: %v \n", p2)
	fmt.Printf("add num: %v \n", p3)
}

func testFunc(num1 int, num2 int) (sum int, reduce int) {
	sum = num1 + num2
	reduce = num1 - num2
	return
}

func sum(n1 int, args ...int) (sum int) {
	sum = n1
	for _, val := range args {
		fmt.Printf("%v\n", val)
		sum += val
	}
	fmt.Printf("%v", sum)
	return
}

func main() {
	sum(1, 2, 3, 4, 5, 6)
}
