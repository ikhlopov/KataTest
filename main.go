package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumType int

const (
	Unknown NumType = iota
	Arabic
	Roman
)

var (
	InvalidOperatorErr  = fmt.Errorf("некорректный оператор")
	InvalidOperandErr   = fmt.Errorf("некорректное число")
	SystemsNotSameErr   = fmt.Errorf("используются одновременно разные системы счисления")
	NegativeRomanResult = fmt.Errorf("в римской системе нет отрицательных чисел")
)

func NumParse(s string) (int, NumType) {

	if num, err := strconv.Atoi(s); err == nil {
		return num, Arabic
	}
	if num, ok := romanDict[s]; ok {
		return num, Roman
	}
	return 0, Unknown
}

var romanDict = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func Count(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		return a / b, nil
	case "*":
		return a * b, nil
	default:
		return 0, InvalidOperatorErr
	}
}

type romNum struct {
	arabic int
	roman  string
}

var ItorNums = []romNum{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Функция для преобразования числа в римскую запись
func Itor(num int) string {
	var roman []string
	for _, r := range ItorNums {
		for num >= r.arabic {
			roman = append(roman, r.roman)
			num -= r.arabic
		}
	}
	return strings.Join(roman, "")
}

func Execute(a, operator, b string) (string, error) {

	numA, typeA := NumParse(a)
	numB, typeB := NumParse(b)

	if typeA == Unknown || typeB == Unknown {
		return "", InvalidOperandErr
	}

	if typeA != typeB {
		return "", SystemsNotSameErr
	}

	isRoman := typeA == Roman

	if numA > 10 || numB > 10 {
		return "", InvalidOperandErr
	}

	if numA < 1 || numB < 1 {
		return "", InvalidOperandErr
	}

	res, err := Count(numA, numB, operator)
	if err != nil {
		return "", err
	}
	if res < 1 && isRoman {
		return "", NegativeRomanResult
	}

	if isRoman {
		return Itor(res), nil
	}

	return strconv.Itoa(res), nil
}

func main() {
	if len(os.Args) != 4 {
		panic("Введите аргументы команды в формате: <число> <оператор> <число>")
		return
	}
	result, err := Execute(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		panic("Ошибка: " + err.Error())
		return
	}
	fmt.Println(result)
}
