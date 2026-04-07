package main

import "fmt"

func main() {
	var (
		a int
	)
	_, _ = fmt.Scan(&a)
	switch {
	case a == 0:
		fmt.Println("ноль")
	case a < 0:
		fmt.Println("Число отрицательное")
	case a > 0:
		fmt.Println("Чосло положительное")
	}
}
