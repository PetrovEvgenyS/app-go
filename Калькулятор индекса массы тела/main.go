package main

import (
	"errors"
	"fmt"
	"math"
)

// Функция для получения данных от пользователя: рост и вес
func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight) // Считываем рост
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg) // Считываем вес
	return userKg, userHeight
}

// Функция для расчета индекса массы тела (IMT)
func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	// Проверяем, что введенные данные корректны
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("!! не указан рост или вес !!")
	}
	const IMTPower = 2
	// Формула расчета IMT: вес (кг) / (рост (м) ^ 2)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

// Функция для вывода результата и интерпретации IMT
func outputResult(imt float64) {
	// Выводим рассчитанный IMT
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)

	// Интерпретация значения IMT
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточная масса тела (предожирение)")
	case imt < 35:
		fmt.Println("У вас ожирение первой степени")
	case imt < 40:
		fmt.Println("У вас ожирение второй степени")
	default:
		fmt.Println("У вас ожирение третьей степени (морбидное)")
	}
}

// Функция для проверки, хочет ли пользователь повторить расчет
func checkRepeateCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать ещё расчёт (y/n): ")
	fmt.Scan(&userChoise) // Считываем выбор пользователя
	// Если пользователь ввел "y" или "Y", продолжаем
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}

// Главная функция программы
func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	for {
		// Получаем данные от пользователя
		userKg, userHeight := getUserInput()
		// Рассчитываем IMT
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			// Если данные некорректны, выводим ошибку и повторяем ввод
			fmt.Println(err)
			continue
		}
		// Выводим результат
		outputResult(IMT)
		// Проверяем, хочет ли пользователь повторить расчет
		isRepeateCalculation := checkRepeateCalculation()
		if !isRepeateCalculation {
			// Если нет, выходим из цикла
			break
		}
	}
}
