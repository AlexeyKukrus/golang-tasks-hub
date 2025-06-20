// Вычисление общей суммы стоимости продуктов
// Наша команда разрабатывает сервис, который должен выполнять следующие действия:

// Принимает две строки, в переменных priceStr и quantityStr. Внутри этих строк записаны числа:

// priceStr — цена продукта (дробное число).
// quantityStr — количество продуктов (целое число).
// Сервис должен вычислить общую сумму, которая равна произведению цены на количество.

// Результат (общую сумму) необходимо вывести в консоль.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	priceStr := "100"
	quantityStr := "5"
	
	priceInt, priceErr := strconv.ParseFloat(priceStr, 64) 
	quantityInt, quantityErr := strconv.ParseFloat(quantityStr, 64)

	if priceErr != nil {
		fmt.Println(priceErr)
	} else if quantityErr != nil {
		fmt.Println(quantityErr)
	} else {
		totalPrice := priceInt * quantityInt 
		fmt.Println(strconv.FormatFloat(float64(totalPrice), 'f', 2, 64))
	}
	
	
}
