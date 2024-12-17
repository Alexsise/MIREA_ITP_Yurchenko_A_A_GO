package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	for {
		fmt.Println("\nВыберите задачу:")
		fmt.Println("1. Линейные задачи")
		fmt.Println("2. Задачи с условным оператором")
		fmt.Println("3. Задачи с циклами")
		fmt.Println("0. Выход")
		fmt.Print("Введите номер категории: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			linearMenu()
		case 2:
			conditionalMenu()
		case 3:
			loopMenu()
		case 0:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Некорректный ввод.")
		}
	}
}

// Меню для линейных задач
func linearMenu() {
	fmt.Println("\nЛинейные задачи:")
	fmt.Println("1. Сумма цифр числа")
	fmt.Println("2. Преобразование температуры")
	fmt.Println("3. Удвоение каждого элемента массива")
	fmt.Println("4. Объединение строк")
	fmt.Println("5. Расчет расстояния между двумя точками")
	fmt.Print("Выберите задачу: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		linearSumDigits()
	case 2:
		linearTemperature()
	case 3:
		linearDoubleArray()
	case 4:
		linearJoinStrings()
	case 5:
		linearDistance()
	default:
		fmt.Println("Некорректный ввод.")
	}
}

// Линейные задачи
func linearSumDigits() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	sum := number%10 + (number/10)%10 + (number/100)%10 + (number/1000)%10
	fmt.Println("Сумма цифр числа:", sum)
}

func linearTemperature() {
	var temp float64
	var scale string
	fmt.Print("Введите температуру и шкалу (Celsius/Fahrenheit): ")
	fmt.Scan(&temp, &scale)

	if strings.EqualFold(scale, "Celsius") {
		fmt.Println("Температура в Фаренгейтах:", temp*1.8+32, "°F")
	} else if strings.EqualFold(scale, "Fahrenheit") {
		fmt.Println("Температура в Цельсиях:", (temp-32)/1.8, "°C")
	} else {
		fmt.Println("Неизвестная шкала.")
	}
}

func linearDoubleArray() {
	var n int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := range arr {
		arr[i] *= 2
	}
	fmt.Println("Удвоенный массив:", arr)
}

func linearJoinStrings() {
	var n int
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&n)

	stringsArray := make([]string, n)
	fmt.Println("Введите строки:")
	for i := 0; i < n; i++ {
		fmt.Scan(&stringsArray[i])
	}

	fmt.Println("Объединенная строка:", strings.Join(stringsArray, " "))
}

func linearDistance() {
	var x1, y1, x2, y2 float64
	fmt.Println("Введите координаты первой точки (x1 y1):")
	fmt.Scan(&x1, &y1)
	fmt.Println("Введите координаты второй точки (x2 y2):")
	fmt.Scan(&x2, &y2)

	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Println("Расстояние между точками:", distance)
}

// Меню для задач с условным оператором
func conditionalMenu() {
	fmt.Println("\nЗадачи с условным оператором:")
	fmt.Println("1. Проверка на четность/нечетность")
	fmt.Println("2. Проверка високосного года")
	fmt.Println("3. Определение наибольшего из трех чисел")
	fmt.Println("4. Категория возраста")
	fmt.Println("5. Проверка делимости на 3 и 5")
	fmt.Print("Выберите задачу: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		conditionalEvenOdd()
	case 2:
		conditionalLeapYear()
	case 3:
		conditionalMaxOfThree()
	case 4:
		conditionalAgeGroup()
	case 5:
		conditionalDivisible()
	default:
		fmt.Println("Некорректный ввод.")
	}
}

// Задачи с условным оператором
func conditionalEvenOdd() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

func conditionalLeapYear() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Невисокосный")
	}
}

func conditionalMaxOfThree() {
	var a, b, c int
	fmt.Print("Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	fmt.Println("Наибольшее число:", max)
}

func conditionalAgeGroup() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	if age >= 0 && age <= 12 {
		fmt.Println("Ребенок")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Подросток")
	} else if age >= 18 && age <= 64 {
		fmt.Println("Взрослый")
	} else if age >= 65 {
		fmt.Println("Пожилой")
	} else {
		fmt.Println("Некорректный возраст")
	}
}

func conditionalDivisible() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}

// Меню для задач с циклами
func loopMenu() {
	fmt.Println("\nЗадачи с циклами:")
	fmt.Println("1. Факториал числа")
	fmt.Println("2. Числа Фибоначчи")
	fmt.Println("3. Реверс массива")
	fmt.Println("4. Простые числа до N")
	fmt.Println("5. Сумма чисел в массиве")
	fmt.Print("Выберите задачу: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		loopFactorial()
	case 2:
		loopFibonacci()
	case 3:
		loopReverseArray()
	case 4:
		loopPrimes()
	case 5:
		loopSumArray()
	default:
		fmt.Println("Некорректный ввод.")
	}
}

// Задачи с циклами
func loopFactorial() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println("Факториал:", fact)
}

func loopFibonacci() {
	var n int
	fmt.Print("Введите количество чисел Фибоначчи: ")
	fmt.Scan(&n)

	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func loopReverseArray() {
	var n int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Перевернутый массив:")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func loopPrimes() {
	var n int
	fmt.Print("Введите число N: ")
	fmt.Scan(&n)

	fmt.Println("Простые числа до", n, ":")
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func loopSumArray() {
	var n int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		sum += arr[i]
	}
	fmt.Println("Сумма элементов массива:", sum)
}
