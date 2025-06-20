// Обработка и анализ строки
// Необходимо выполнить следующие действия с данной строкой:

// Удалить пробелы слева и справа от строки message, если такие имеются, и вывести полученную строку.
// Преобразовать строку, полученную в пункте 1, в нижний регистр и вывести её.
// Вывести true, если строка, полученная в пункте 2, начинается с букв "go", в противном случае вывести false.
// Весь вывод должен осуществляться с помощью стандартных функций, вроде fmt.Println().

package main

import (
	"strings"
	"fmt"
	
)

func main() {
	message := " Go - это не просто язык, это СТИЛЬ ЖИЗНИ! "
	trimmedStr := strings.Trim(message, " ")
	lowerStr := strings.ToLower(trimmedStr)
	hasPrefix := strings.HasPrefix(lowerStr, "go")
	fmt.Println(trimmedStr)
	fmt.Println(lowerStr)
	fmt.Println(hasPrefix)
}