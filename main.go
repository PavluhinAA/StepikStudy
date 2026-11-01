package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
)

type Items struct {
	GlobalID int64 `json:"global_id"`
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if ferr := file.Close(); ferr != nil {
			log.Println(ferr)
		}
	}()

	input, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var data []Items
	err = json.Unmarshal(input, &data)
	if err != nil {
		log.Fatal(err)
	}

	var sum int64
	for _, item := range data {
		sum += item.GlobalID
	}

	_, err = os.Stdout.WriteString(strconv.FormatInt(sum, 10) + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
