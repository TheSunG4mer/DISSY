package main

import (
	"os"
)

func main() {
	fi, err := os.Open("C:/Users/au649790/OneDrive - Aarhus universitet/Desktop/DISSY/week5/test.txt")
	if err != nil {
		panic(err)
	}

	fi.Close()
}
