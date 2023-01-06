package main

import "fmt"

func main() {
	sampleArray := []int{1, 2, 2, 1, 4, 5, 7, 5, 6, 7, 6}
	unique := findUnique(sampleArray)
	printUniqueNum(unique)
}

func findUnique(numArr []int) int {
	numMap := make(map[int]struct{})
	for _, num := range numArr {
		if _, ok := numMap[num]; !ok {
			numMap[num] = struct{}{}
		} else {
			delete(numMap, num)
		}
	}

	for num, _ := range numMap {
		return num
	}
	return 0
}

func printUniqueNum(num int) {
	fmt.Println("the unique number is :", num)
}
