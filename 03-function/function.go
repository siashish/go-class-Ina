package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// multiple return
func swap(x, y string) (string, string) {
	return y, x
}

// named return
func split(sum int) (x int, y float32) {
	x = sum * 4 / 9
	y = float32(sum - x)
	return
}

func split1(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))

	a, b := swap("hello", "Ina")
	fmt.Println(a, b)

	fmt.Println(split(17))
	fmt.Println(split1(17))
}
