// Получение информации о пользователе
// Вам необходимо реализовать функцию userProfile, которая будет обрабатывать информацию о пользователе на основе его идентификатора.

// Описание функции
// Функция userProfile возвращает информацию о пользователе.

// Параметры:

// id (тип string) — идентификатор пользователя.
// Возвращаемые значения:

// string — информация о пользователе.
// error — ошибка, если она произошла.
// Алгоритм работы функции
// Внутри функции userProfile вызовите уже реализованную функцию fetchUserInfo(id string) (int, error), которая принимает идентификатор пользователя, а возвращает его баланс (в копейках) и ошибку, если таковая имеется.

// Если fetchUserInfo вернула ошибку, то верните ошибку и из userProfile, обернув её в строку: "fetch error: [ОШИБКА_ИЗ_fetchUserInfo]".

// Если fetchUserInfo вернула данные без ошибки, то необходимо выполнить следующие действия:
// Переведите баланс (в копейках) в рубли с копейками (тип float64).
// Верните сообщение в формате: "Пользователь с id [ID_ПОЛЬЗОВАТЕЛЯ] имеет на счету [БАЛАНС] руб.".

package main

import (
	"fmt"
	"errors"
)

func userProfile(id string) (string, error) {
	userInfo, err := fetchUserInfo(id)
	if err != nil {
		fmt.Errorf("fetch error: %w", err)
	}
	return fmt.Sprintf("Пользователь с id %v имеет на счету %v руб.", id, userInfo), nil
}

func fetchUserInfo(id string) (int, error) {
	return 0, errors.New("error mf")
}

func main() {
	fmt.Println(userProfile("1234567"))
}