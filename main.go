package main

import "fmt"

func main() {
	workArray := [10]uint8{}
	for i := 0; i < len(workArray); i++ {
		fmt.Scanf("%d", &workArray[i])
	}

	for i := 0; i < 3; i++ {
		var ind1, ind2 int
		fmt.Scan(&ind1, &ind2)
		workArray[ind1], workArray[ind2] = workArray[ind2], workArray[ind1]

	}
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}
