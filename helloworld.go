package main

import (
	"fmt"
)

func main() {

	var price int = 100
	fmt.Println("Price is", price, "dollars")

	var taxRate float64 = 0.08
	var tax float64 = taxRate * float64(price)

	fmt.Println("Tax is", tax, "dollars")
}
