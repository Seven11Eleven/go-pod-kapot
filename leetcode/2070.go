package leetcode

import "fmt"

func MaximumBeauty(items [][]int, queries []int) []int {

	items = [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
	queries = []int{1, 2, 3, 4, 5, 6}
	beatuies := make(map[int]int)
	res := make([]int, 0)
	for _, val := range items {
		//for _, _ := range val {
		//	beatuies[val[1]] = val[0]
		//}
		beatuies[val[0]] = val[1]

		fmt.Printf("i: %d\nval: %d\n", val[0], val[1])
	}
	//for i, val := range queries {
	//	if beatuies[val] > -1 {
	//		if beatuies[val] <= val {
	//			res = append(res, beatuies[val])
	//		}
	//	}
	//	fmt.Printf("i: %d\nval;: %d\n", i, val)
	//}
	return res
}
