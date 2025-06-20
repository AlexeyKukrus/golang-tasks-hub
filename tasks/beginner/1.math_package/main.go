package main

import (
	"fmt"
	"math"
)

func main() {
	var discountPercent float64 = 11.111
	var productPrice float64 = 100.0

	discount := productPrice * discountPercent / 100
	finalPrice := productPrice - discount
	roundedPrice := math.Floor(finalPrice*100) / 100
	fmt.Println(roundedPrice)
}