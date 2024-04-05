package main

import "fmt"

func main() {
	// fmt.Println(twoSum([]int{1, 4, 5, 7}, 8))
	fmt.Println(twoSum2([]int{1, 4, 5, 7}, 8))
}

// Dever retornar os indices do array em que a soma seja igual ao alvo.
// ex: list[1 4 5 7] -> 8
//
//	1 + 7 -> index "0" e index "3"
func twoSum(nums []int, target int) []int {
	var indexs []int

	for i := range nums {

		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indexs = []int{i, j}
				break
			}
		}

		if len(indexs) > 0 {
			break
		}
	}

	return indexs
}

func twoSum2(nums []int, target int) []int {
	numerosCandidatos := make(map[int]int)

	for index, num := range nums {
		i, exist := numerosCandidatos[num]
		if exist {
			return []int{i, index}
		}

		numerosCandidatos[target-num] = index
	}

	return []int{}
}
