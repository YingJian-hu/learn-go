package main

import "fmt"

func Delete[T any](idx int, vals []T) []T {
	if idx < 0 || idx > len(vals)-1 {
		panic("error")
	}

	left := vals[:idx]
	right := vals[idx+1:]
	vals = append(left, right...)

	// 通过新建一个数组的方式进行缩容，重开空间，然后把结果copy过来
	// 将老空间释放掉
	var res []T
	for _, val := range vals {
		res = append(res, val)
	}

	return res
}

func SliceDelete() {
	// int
	sInt := []int{1, 2, 3}
	fmt.Printf("删除前, len: %d, cap: %d \n", len(sInt), cap(sInt)) // cop = 3
	fmt.Printf("sInt: %v \n", Delete(1, sInt))
	fmt.Printf("删除后, len: %d, cap: %d \n", len(Delete(1, sInt)), cap(Delete(1, sInt))) // cop = 2
	fmt.Printf("\n")

	// map
	sMap := []map[string]int{
		map[string]int{
			"age": 18,
		},
		map[string]int{
			"age": 20,
		},
		map[string]int{
			"age": 22,
		},
	}
	fmt.Printf("删除前, len: %d, cap: %d \n", len(sMap), cap(sMap))
	fmt.Printf("sInt: %v \n", Delete(1, sMap))
	fmt.Printf("删除后, len: %d, cap: %d \n", len(Delete(1, sMap)), cap(Delete(1, sMap)))

}
