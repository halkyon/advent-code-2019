package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	opAdd  = 1
	opMult = 2
	opHalt = 99
)

func main() {
	nums := make(map[int]int)
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for i, num := range strings.Split(string(input), ",") {
		num, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = num
	}

	// part 1
	num := calculate(nums[1], nums[2], nums)
	fmt.Println(num)

	// part 2
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if calculate(i, j, nums) == 19690720 {
				fmt.Println(100*i + j)
				break
			}
		}
	}
}

func calculate(num1 int, num2 int, nums map[int]int) int {
	// make a copy of the map so we don't cause side-effects to the original
	tmp := make(map[int]int)
	for i, v := range nums {
		tmp[i] = v
	}

	tmp[1] = num1
	tmp[2] = num2
	idx := 0
	for tmp[idx] != 99 {
		num := tmp[idx]
		src1 := tmp[tmp[idx+1]]
		src2 := tmp[tmp[idx+2]]
		destIdx := tmp[idx+3]
		switch num {
		case opAdd:
			tmp[destIdx] = src1 + src2
		case opMult:
			tmp[destIdx] = src1 * src2
		}
		idx += 4
	}
	return tmp[0]
}
