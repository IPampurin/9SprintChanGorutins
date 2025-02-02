package main

import (
	"fmt"
	"time"
)

func main() {
	// напишите код программы
	ticker := time.NewTicker(200 * time.Millisecond)

	for i := 1; i <= 20; i++ {
		<-ticker.C
		if i%5 == 0 {
			fmt.Print("o")
			continue
		}
		fmt.Print(".")
	}
}
