package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("./misc/commits.csv")
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}
