package main

import (
	"fmt"
	"os"
)

func main() {
	value1, value2, operation := readTask()
	if v, ok := value1.(float64); ok {
		if v2, ok2 := value2.(float64); ok2 {
			if str, ok3 := operation.(string); ok3 {
				switch str {
				case "+":
					fmt.Printf("%.4f", v2+v)
					os.Exit(0)
				case "-":
					fmt.Printf("%.4f", -v2+v)
					os.Exit(0)
				case "*":
					fmt.Printf("%.4f", v2*v)
					os.Exit(0)
				case "/":
					fmt.Printf("%.4f", v/v2)
					os.Exit(0)
				default:
					fmt.Println("неизвестная операция")
					os.Exit(0)
				}
			}
			fmt.Println("неизвестная операция")
			os.Exit(0)
		}
		fmt.Printf("value=%v: %T\n", value2, value2)
		os.Exit(0)
	}
	fmt.Printf("value=%v: %T\n", value1, value1)
	os.Exit(0)
}
func readTask() (interface{}, interface{}, interface{}) {
	var (
		a = 5.6
		b = 3.2
		c = 3
	)
	return a, b, c
}
