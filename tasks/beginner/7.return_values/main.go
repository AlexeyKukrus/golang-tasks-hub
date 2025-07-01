package main

import (
	"fmt"
	"math/rand"
)
var name string = "Павел"
func generateCompliment(name string) string {
	var random = rand.Intn(3)
	var result string

	switch random {
	case 0:
		result = fmt.Sprintf("Ты великолепен, %v!", name)
	case 1:
		result = fmt.Sprintf("У тебя потрясающая улыбка, %v!", name)
	case 2: 
		result = fmt.Sprintf("Ты вдохновляешь, %v!", name)
	default:
		result = ""
	}
	return result
}

func main() {
	generateCompliment(name)
}