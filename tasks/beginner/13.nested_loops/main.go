// 1. Таблица умножения
// Хочется вывести таблицу умножения для ребенка. Необходимо написать функцию printTable, которая принимает число (назовем параметр num) и выводит 
// таблицу умножения num x num.
// Используйте вложенные циклы и форматируйте вывод в виде таблицы, между примерами отступ должен состоять из одного знака табуляции, в конце каждой 
// строки табуляции быть не должно.

// 2. Ромб из звездочек
// Порисуем? Хочется увидеть ромбик в терминале, для этого необходимо создать функцию func printDiamond(n int), которая выводит текст Мой бриллиант:, 
// а с новой строки ромб (смотри примеры), где n - это ребро ромба.
// Мой бриллиант:
//  #
// # #
//  #

package main

import (
	"fmt"
)

func printTable(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			if j == num {
				fmt.Printf("%v x %v = %v \n", i, j, i*j)
			} else {
				fmt.Printf("%v x %v = %v \t", i, j, i*j)
			}
			

		}
	}
}

func printDiamond(num int) {
    fmt.Println("Мой бриллиант:")

    for i := 0; i < num; i++ {
        for j := 0; j < num-i-1; j++ {
            fmt.Print(" ")
        }
        fmt.Print("#")
        if i > 0 {
            for j := 0; j < 2*i-1; j++ {
                fmt.Print(" ")
            }
            fmt.Print("#")
        }
        fmt.Println()
    }

    for i := num - 2; i >= 0; i-- {
        for j := 0; j < num-i-1; j++ {
            fmt.Print(" ")
        }
        fmt.Print("#")
        if i > 0 {
            for j := 0; j < 2*i-1; j++ {
                fmt.Print(" ")
            }
            fmt.Print("#")
        }
        fmt.Println()
    }
}

func main() {
	printTable(4)
	printDiamond(4)
}