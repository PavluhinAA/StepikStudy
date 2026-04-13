package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(recursTest(a))
}
func recursTest(a int) int { //Чистый тест, никакого практического применения не имеет
	if a < 10 {
		return a
	}
	str := strconv.Itoa(a)
	arr := strings.Split(str, "")
	a = 0
	for i := 0; i < len(arr); i++ {
		b, _ := strconv.Atoi(arr[i])
		a += b
	}
	return recursTest(a)
}
