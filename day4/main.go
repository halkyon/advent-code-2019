package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	count := 0
	for password := 278384; password <= 824795; password++ {
		if !ascendingDigits(password) {
			continue
		}
		if !twoDigitsSame(password) {
			continue
		}
		count++
	}
	fmt.Println(count)
}

func ascendingDigits(password int) bool {
	tmp := strconv.Itoa(password)
	for i := 0; i < 5; i++ {
		if tmp[i] > tmp[i+1] {
			return false
		}
	}
	return true
}

func twoDigitsSame(password int) bool {
	tmp := strconv.Itoa(password)
	for i := 0; i < 5; i++ {
		count := strings.Count(tmp, string(tmp[i]))
		if count == 2 {
			return true
		}
	}
	return false
}
