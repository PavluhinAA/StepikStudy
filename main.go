package main

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const layout = "2006-01-02 15:04:05"

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	tStr := strings.TrimSpace(string(input))

	t, err := time.Parse(layout, tStr)
	if err != nil {
		log.Fatal(err)
	}

	if t.Hour() > 13 {
		t = t.Add(time.Hour * 24)
	}

	_, err = os.Stdout.WriteString(t.Format(layout) + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
