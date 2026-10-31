package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Student struct{ Rating []int }
type Group struct{ Students []Student }
type Result struct{ Average float64 }

func main() {
	var (
		data   Group
		numRat int
	)
	/*file, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	input, err := io.ReadAll(file)*/
	input, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(input, &data)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, student := range data.Students {
		numRat += len(student.Rating)
	}
	numStd := len(data.Students)

	if numStd == 0 {
		log.Println("No students found")
		os.Exit(1)
	}

	out, err := json.MarshalIndent(Result{Average: float64(numRat) / float64(numStd)}, "", "    ")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	_, err = os.Stdout.Write(out)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
