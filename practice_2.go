package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("\nЗадачник на Go")
		fmt.Println("1. Линейное программирование")
		fmt.Println("2. Условия")
		fmt.Println("3. Циклы")
		fmt.Println("0. Выход")
		fmt.Print("Введите номер категории: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			linearTasksMenu()
		case 2:
			conditionalTasksMenu()
		case 3:
			loopTasksMenu()
		case 0:
			fmt.Println("Выход из задачника.")
			return
		default:
			fmt.Println("Некорректный ввод.")
		}
	}
}

// Меню для линейных задач
func linearTasksMenu() {
	fmt.Println("\nЛинейное программирование")
	fmt.Println("1. Перевод чисел из одной системы счисления в другую")
	fmt.Println("2. Решение квадратного уравнения")
	fmt.Println("3. Сортировка чисел по модулю")
	fmt.Println("4. Слияние двух отсортированных массивов")
	fmt.Println("5. Нахождение подстроки в строке")
	fmt.Print("Выберите задачу: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		linearConvertBase()
	case 2:
		linearSolveQuadratic()
	case 3:
		linearSortByAbs()
	case 4:
		linearMergeSortedArrays()
	case 5:
		linearFindSubstring()
	default:
		fmt.Println("Некорректный ввод.")
	}
}

// Линейные задачи
func linearConvertBase() {
	var number string
	var fromBase, toBase int
	fmt.Print("Введите число, исходную систему и конечную систему счисления: ")
	fmt.Scan(&number, &fromBase, &toBase)

	n, _ := strconv.ParseInt(number, fromBase, 64)
	result := strconv.FormatInt(n, toBase)
	fmt.Println("Результат:", result)
}

func linearSolveQuadratic() {
	var a, b, c float64
	fmt.Print("Введите коэффициенты a, b, c: ")
	fmt.Scan(&a, &b, &c)

	d := b*b - 4*a*c
	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println("Корни:", x1, x2)
	} else if d == 0 {
		x := -b / (2 * a)
		fmt.Println("Корень:", x)
	} else {
		re := -b / (2 * a)
		im := math.Sqrt(-d) / (2 * a)
		fmt.Printf("Комплексные корни: %.2f+%.2fi, %.2f-%.2fi\n", re, im, re, im)
	}
}

func linearSortByAbs() {
	var n int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})

	fmt.Println("Отсортированный массив:", arr)
}

func linearMergeSortedArrays() {
	var n1, n2 int
	fmt.Print("Введите размер первого массива: ")
	fmt.Scan(&n1)
	fmt.Println("Введите элементы первого массива:")
	a1 := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&a1[i])
	}

	fmt.Print("Введите размер второго массива: ")
	fmt.Scan(&n2)
	fmt.Println("Введите элементы второго массива:")
	a2 := make([]int, n2)
	for i := 0; i < n2; i++ {
		fmt.Scan(&a2[i])
	}

	merged := mergeSortedArrays(a1, a2)
	fmt.Println("Объединенный массив:", merged)
}

func mergeSortedArrays(a1, a2 []int) []int {
	result := make([]int, 0, len(a1)+len(a2))
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			result = append(result, a1[i])
			i++
		} else {
			result = append(result, a2[j])
			j++
		}
	}

	result = append(result, a1[i:]...)
	result = append(result, a2[j:]...)
	return result
}

func linearFindSubstring() {
	var s, sub string
	fmt.Print("Введите строку и подстроку: ")
	fmt.Scan(&s, &sub)

	index := strings.Index(s, sub)
	if index == -1 {
		fmt.Println("Подстрока не найдена.")
	} else {
		fmt.Println("Индекс первого вхождения:", index)
	}
}

func conditionalTasksMenu() {
	for {
		fmt.Println("\nВыберите задачу:")
		fmt.Println("1. Калькулятор с расширенными операциями")
		fmt.Println("2. Проверка палиндрома")
		fmt.Println("3. Нахождение пересечения трех отрезков")
		fmt.Println("4. Выбор самого длинного слова в предложении")
		fmt.Println("5. Проверка високосного года")
		fmt.Println("6. Вернуться в главное меню")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			advancedCalculator()
		case 2:
			checkPalindrome()
		case 3:
			checkIntersection()
		case 4:
			findLongestWord()
		case 5:
			checkLeapYear()
		case 6:
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func advancedCalculator() {
	var a, b float64
	var operator string
	fmt.Print("Введите два числа и оператор (+, -, *, /, ^, %): ")
	fmt.Scanln(&a, &b, &operator)

	switch operator {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Результат:", a/b)
		} else {
			fmt.Println("Ошибка: деление на ноль!")
		}
	case "^":
		fmt.Println("Результат:", math.Pow(a, b))
	case "%":
		if b != 0 {
			fmt.Println("Результат:", math.Mod(a, b))
		} else {
			fmt.Println("Ошибка: деление на ноль!")
		}
	default:
		fmt.Println("Недопустимая операция.")
	}
}

func checkPalindrome() {
	var input string
	fmt.Print("Введите строку для проверки палиндрома: ")
	fmt.Scanln(&input)

	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	isPalindrome := true
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}

func checkIntersection() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Print("Введите отрезки (x1 y1 x2 y2 x3 y3): ")
	fmt.Scanln(&x1, &y1, &x2, &y2, &x3, &y3)

	intersectionStart := math.Max(x1, math.Max(x2, x3))
	intersectionEnd := math.Min(y1, math.Min(y2, y3))

	if intersectionStart <= intersectionEnd {
		fmt.Println("Отрезки пересекаются.")
	} else {
		fmt.Println("Отрезки не пересекаются.")
	}
}

func findLongestWord() {
	var sentence string
	fmt.Print("Введите предложение: ")
	fmt.Scanln(&sentence)

	words := strings.FieldsFunc(sentence, func(r rune) bool {
		return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || '0' <= r && r <= '9')
	})

	longestWord := ""
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	fmt.Println("Самое длинное слово:", longestWord)
}

func checkLeapYear() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Год високосный.")
	} else {
		fmt.Println("Год не високосный.")
	}
}

func loopTasksMenu() {
	for {
		fmt.Println("\nВыберите задачу:")
		fmt.Println("1. Числа Фибоначчи до определенного числа")
		fmt.Println("2. Определение простых чисел в диапазоне")
		fmt.Println("3. Числа Армстронга в заданном диапазоне")
		fmt.Println("4. Реверс строки")
		fmt.Println("5. Нахождение наибольшего общего делителя (НОД)")
		fmt.Println("6. Вернуться в главное меню")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fibonacciUntilN()
		case 2:
			findPrimesInRange()
		case 3:
			findArmstrongNumbers()
		case 4:
			reverseString()
		case 5:
			findGCD()
		case 6:
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func fibonacciUntilN() {
	var n int
	fmt.Print("Введите максимальное число для последовательности Фибоначчи: ")
	fmt.Scanln(&n)

	a, b := 0, 1
	fmt.Print("Последовательность Фибоначчи: ")
	for a <= n {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func findPrimesInRange() {
	var start, end int
	fmt.Print("Введите начало и конец диапазона: ")
	fmt.Scanln(&start, &end)

	fmt.Print("Простые числа: ")
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findArmstrongNumbers() {
	var start, end int
	fmt.Print("Введите начало и конец диапазона: ")
	fmt.Scanln(&start, &end)

	fmt.Print("Числа Армстронга: ")
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func isArmstrong(num int) bool {
	strNum := strconv.Itoa(num)
	power := len(strNum)
	sum := 0
	for _, digit := range strNum {
		d, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(d), float64(power)))
	}
	return sum == num
}

func reverseString() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	reversed := ""
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	fmt.Println("Реверс строки:", reversed)
}

func findGCD() {
	var a, b int
	fmt.Print("Введите два числа: ")
	fmt.Scanln(&a, &b)

	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println("Наибольший общий делитель:", a)
}
