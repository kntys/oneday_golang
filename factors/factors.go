package main

import "fmt"

func getFactors(num int) {
	fmt.Println(num)
	for n := 0; n < num; n++ {
		switch {
		case n == 0:
		case n == 1:
		case num%n == 0:
			num /= n
			fmt.Printf("を%dで割ると%d \n", n, num)
		default:

		}
	}
	fmt.Println()
}

func main() {
	getFactors(2310)
	getFactors(37789)
	getFactors(1256997)
}
