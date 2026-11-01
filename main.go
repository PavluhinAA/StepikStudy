package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const layout = "02.01.2006 15:04:05"

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	twoTime := strings.Split(strings.TrimSpace(string(input)), ",")
	if len(twoTime) != 2 {
		log.Fatal("invalid input")
	}
	t1, err := time.Parse(layout, twoTime[0])
	if err != nil {
		log.Fatal("t1", err)
	}
	t2, err := time.Parse(layout, twoTime[1])
	if err != nil {
		log.Fatal("t2", err)
	}

	if t2.After(t1) {
		t1, t2 = t2, t1
	}

	fmt.Println(t1.Sub(t2))
}
