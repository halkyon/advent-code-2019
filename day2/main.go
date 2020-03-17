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
	var nums []int
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, num := range strings.Split(string(input), ",") {
		num, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	for i, num := range nums {
		if i%4 != 0 {
			continue
		}
		src1 := nums[nums[i+1]]
		src2 := nums[nums[i+2]]
		destIdx := nums[i+3]
		switch num {
		case opAdd:
			nums[destIdx] = src1 + src2
		case opMult:
			nums[destIdx] = src1 * src2
		case opHalt:
			fmt.Printf("%+v\n", nums)
			return
		default:
			log.Fatalf("cannot handle opcode %d\n", num)
		}
	}
}
