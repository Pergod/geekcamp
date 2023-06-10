package main

import (
	"fmt"
)

func main() {
	array := []string{"I", "am", "stupid", "and", "weak"}
	changeStringArray(array)
	fmt.Println(array)
}

func changeStringArray(arr []string) {
	for index, _ := range arr {
		if arr[index] == "stupid" {
			arr[index] = "smart"
		} else if arr[index] == "weak" {
			arr[index] = "strong"
		}
	}
}
