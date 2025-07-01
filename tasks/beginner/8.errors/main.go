// 1. Напишите функцию UserProfileToString, которая принимает имя пользователя name (string) и его возраст age (int).
// Функция должна возвращать строку с сообщением (string) и ошибку (error), если таковая была.
// Сообщение должно быть в формате: Имя человека: [ИМЯ], возраст: [ВОЗРАСТ].
// Если имя было передано в функцию и имеет пробелы слева/справа - эти пробелы нужно убрать.
// Возможные ошибки
// Если имя пользователя пустое, функция должна вернуть ошибку с сообщением: empty name
// Если возраст меньше нуля, необходимо вернуть ошибку с сообщением: negative age
// Если передано имя, состоящее только из пробелов, то функция должна вернуть ошибку с сообщением: name cannot contain only spaces 
// Примечание
// Создавать функцию main не нужно, она уже создана за вас и уже вызывает функцию UserProfileToString и проверяет корректность ее работы.
// В этом задании ничего выводить не нужно, любой вывод будет считаться за ошибку.

// 2. Реализуйте функцию calculate, которая принимает три аргумента:
// Число (тип float64) - первый аргумент для вычисления.
// Число (тип float64) - второй аргумент для вычисления.
// Строку, представляющую операцию, которую необходимо выполнить. Возможные операции:
// "add" (сложение)
// "subtract" (вычитание)
// "multiply" (умножение)
// "divide" (деление)
// Функция должна возвращать результат операции и ошибку.
// Ошибки
// Если операция не поддерживается, функция должна возвращать ошибку с сообщением: unknown operation.
// Если происходит деление на ноль, возвращаем ошибку с сообщением division by zero.
// В случае любой ошибки, первое возвращаемое значение должно быть равно 0.

package main

import (
	"fmt"
	"errors"
	"strings"
)

func UserProfileToString(name string, age int) (string, error) {
	var trimmedName = strings.TrimSpace(name)
	if (name == "") {
		return "", errors.New("empty name")
	}
	if (age < 0) {
		return "", errors.New("negative age")
	}
	if (trimmedName == "") {
		return "", errors.New("name cannot contain only spaces")
	}
	return fmt.Sprintf("Имя человека: %v, возраст: %v.", trimmedName, age), nil
}

func calculate(a, b float64, op string) (float64, error) {
    switch op {
    case "add":
        return a + b, nil
    case "subtract":
        return a - b, nil
    case "multiply":
        return a * b, nil
    case "divide":
        if b == 0 {
            return 0, errors.New("division by zero")
        }
        return a / b, nil
    default:
        return 0, errors.New("unknown operation")
    }
}

func main() {
	UserProfileToString("   ", 19)
	calculate(10, 5, "add")
}