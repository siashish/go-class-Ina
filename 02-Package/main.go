package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My fav nunber is", rand.Intn(10))
	fmt.Printf("Now you have %g problems\n", math.Sqrt(4))
	// exported name
	fmt.Println(math.Pi)
}
