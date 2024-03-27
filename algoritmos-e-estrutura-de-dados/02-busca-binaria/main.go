package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 7, 9, 23, 34, 43, 65, 76, 90, 656}
	loops, ok := find(656, list)
	fmt.Println(loops, ok)
}

func find(target int, list []int) (int, bool) {

	var (
		start = 0
		end   = len(list) - 1
		loops = 0
	)

	for {
		loops++
		middleIndex := (start + end) / 2
		if start > end {
			return loops, false
		}

		middleItem := list[middleIndex]

		switch {
		case middleItem == target:
			return loops, true

		case middleItem > target:
			end = middleIndex - 1

		case middleItem < target:
			start = middleIndex + 1
		}
	}

}
