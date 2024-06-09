package main

import (
	"fmt"
	"os"
	"strconv"
)

func romToara(a string) int {

	rml := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	for k, v := range rml {
		if a == k {
			return v
		}
	}
	panic("неизвестное число")

}
func intToRom(n int) string {
	rml := []struct {
		val int
		rom string
	}{
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	result := ""
	for _, v := range rml {
		for n >= v.val {
			result += v.rom
			n -= v.val
		}
	}
	return result
}
func main() {
	var a, b string
	var s string
	var err = ""
	fmt.Println("Введите математическое выражение:")
	fmt.Fscanln(os.Stdin, &a, &s, &b, &err)
	if err != "" {
		fmt.Println("проверьте правильность ввода")

	}
	if _, err := strconv.Atoi(a); err == nil {
		if _, err := strconv.Atoi(b); err != nil {
			fmt.Println("Ошибка: строка b не состоит из цифр")
			return
		}
	} else {
		if _, err := strconv.Atoi(b); err == nil {
			fmt.Println("Ошибка: строка a не состоит из цифр")
			return
		}
	}
	var count int = 0
	numA, _ := strconv.Atoi(a)
	if numA == 0 {
		numA = romToara(a)
		count++
	}
	numB, _ := strconv.Atoi(b)
	if numB == 0 {
		numB = romToara(b)
		count++
	}

	if numA < 1 || numA > 10 {
		fmt.Println("Ошибка: число а не входит в диапазон от 1 до 10")
		return
	}

	if numB < 1 || numB > 10 {
		fmt.Println("Ошибка: число б не входит в диапазон от 1 до 10")
		return
	}
	var result int
	switch s {
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "*":
		result = numA * numB
	case "/":
		result = numA / numB
	default:
		fmt.Println("неизвестный операнд")

	}
	if count == 2 {
		fmt.Println(intToRom(result))
	} else {
		fmt.Println(result)
	}

}
