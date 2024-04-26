package main

import "fmt"

func main() {
	fmt.Println(computerDeclension(34))
	fmt.Println(computerDeclension(5))
	fmt.Println(computerDeclension(17))
}

func computerDeclension(num int) string {
	if num < 0 {
		return "Ошибка: отрицательное число"
	}

	switch {
	case num%100 >= 11 && num%100 <= 20:
		return fmt.Sprintf("%d компьютеров", num)
	case num%10 == 1:
		return fmt.Sprintf("%d компьютер", num)
	case num%10 >= 2 && num%10 <= 4:
		return fmt.Sprintf("%d компьютера", num)
	default:
		return fmt.Sprintf("%d компьютеров", num)
	}
}