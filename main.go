package main

import (
	"fmt"
)

func main() {

	storms := ReadStorms()
	radars := ReadRadars()

	fmt.Println(storms)
	fmt.Println(radars)
}
