package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 3, 6, 3}
	var result []int

	// ініціюємо мапу для отримання уникальних значень
	values := make(map[int]bool)
	
	for _,value := range arr {           
		if _, ok := values[value]; !ok { 
			values[value] = true
			result = append(result, value)
		}
	}
	fmt.Println(result)
}