package main

import (
	"fmt"
	"time"
)

const Count = 1000
const Lines = 4

var data [Lines][Count]int

func main() {
	// вставьте сюда код с запуском четырёх горутин
	for i := 0; i < Lines; i++ {
		go func(strNum int) {
			for j := 0; j < Count; j++ {
				data[strNum][j] = j
			}
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	// проверим как заполнен массив
	sum := 0
	for i := 0; i < Lines; i++ {
		for j := 0; j < Count; j++ {
			sum += data[i][j]
		}
	}
	fmt.Println(sum)

	for i := 0; i < 4; i++ {
		for j := 0; j < 30; j++ {
			fmt.Printf("%4d", data[i][j])
		}
		fmt.Println()
	}
}
