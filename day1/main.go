package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err.Error())
		}
		sum += calculateFuel(value)
	}
	fmt.Println(sum)
}

func calculateFuel(value int) int {
	fuel := value/3 - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + calculateFuel(fuel)
}
