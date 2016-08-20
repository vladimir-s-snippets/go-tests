package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Printf("%d: %s\n", i, arg)
	}

	s1, sep := "", " "
	data := make([]int, 100000)
	start1 := time.Now()
	for _, val := range data {
		//строка пересоздаётся, вызывается сборщик мусора
		s1 += strconv.Itoa(val) + sep
	}

	duration1 := -1 * start1.Sub(time.Now())
	fmt.Println(duration1.Nanoseconds())

	start2 := time.Now()
	dataStr := make([]string, 100000)
	for i, val := range data {
		dataStr[i] = strconv.Itoa(val)
	}
	//этот метод не пересоздаёт строку
	strings.Join(dataStr, sep)

	duration2 := -1 * start2.Sub(time.Now())
	fmt.Println(duration2.Nanoseconds())
}
