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
	for _, v := range arrString {
		x, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		arr[i] = x
		i++
	}

	var minValue = math.MaxInt32
	for _, v := range arr {
		if v < minValue {
			minValue = v
		}
	}

	var maxValue = math.MinInt32
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}

	result = fmt.Sprint(maxValue, " та ", minValue)
	fmt.Println(result)
}
