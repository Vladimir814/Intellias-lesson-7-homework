package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	 var result string

	 arrString := strings.Split(input, " ")
	 arr := make([]int, len(arrString))
	 i := 0
	 for _,v := range arrString {
		x,err := strconv.Atoi(v)
	    if err != nil {
		    panic(err)
	    }
	    arr[i] = x
		i++
	}

	var minValue = min(arr)
	var maxValue = max(arr)
	result = fmt.Sprint(maxValue, " та ", minValue)
	fmt.Println(result)
}

func min(arr []int) int {
	var res = math.MaxInt32
	for _, v := range arr {
		if v < res {
			res = v
		}
	}
	return res
}

func max(arr []int) int {
	var res = math.MinInt32
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}
