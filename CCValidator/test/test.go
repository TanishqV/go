package main

import (
	"fmt"
	"os"
	"math/rand"
	"strings"
	"time"
	"CCValidator/ccvalidator"
)

// Random String generation using: https://stackoverflow.com/a/31832326
const letterBytes = "12345678909876543210"
const (
	letterIdxBits = 6 // 6 bits to represent letter index
	letterIdxMask = 1 << letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax = 63/letterIdxBits // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randStringGenerator(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i,cache,remain := n-1,src.Int63(),letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func generateNums(part string) (res []string) {
	lens := [...]int {13,14,15,16}
	for _,v :=range lens {
		num := randStringGenerator(v - len(part))
		num = part + num
		res = append(res, num)
	}
	return
}

func validateNums(nums []string) {
	for _,num := range nums {
		fmt.Print(num)
		manufacturer, ok, err := ccvalidator.Validate(num)
		if ok {
				fmt.Printf("||\tManufacturer: %s", manufacturer)
		} else {
				fmt.Printf("||\t%s", err)
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) == 3 &&
	   len(os.Args[1]) >= 5 && len(os.Args[1]) <= 6 &&
	   len(os.Args[2]) >= 5 && len(os.Args[2]) <= 6 {
		var nums []string
		for i:=1; i<len(os.Args); i++ {
			fmt.Printf("For %v,", os.Args[i])
			nums = generateNums(os.Args[i])
			fmt.Println()
			validateNums(nums)
			fmt.Println()
		}
	} else {
		fmt.Println("Enter 2 numbers of 5 or 6 digits each as the command line arguments\n")
	}
}
