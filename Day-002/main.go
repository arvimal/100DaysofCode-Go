package main

import "fmt"

const PI = 3.1412
const B string = "Hello"
const BEEF, TWO, C = "meat", 2, "see"

const (
	UNKNOWN = 0
	FEMALE  = 1
	MALE    = 2
)

var bingo = 10

func main() {
	var number float32 = 5.2
	fmt.Println(number)
	fmt.Println(int(number))
}
