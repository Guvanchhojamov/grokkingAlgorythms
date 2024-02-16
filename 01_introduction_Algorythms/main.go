package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Binary search
// 1) Given sorted list with 128 item, and we search [word] in them, with method binary search,
// well, what count of maximum operations we need for find the [word] from this list?

// Response:
// we need: max=8 operations
//
// Explanation:
// n:2 in every operation gives this
// 128->64->32->16->8->4->2->1  = 8

// 2) just response: O(2*log2n)

// Phrases
// linear time: O(n)
// logaryphmik time O(log2n)

//Типичные примеры «О-большого»
//Ниже перечислены пять разновидностей «О-большого», которые будут встре
//чаться вам особенно часто, в порядке убывания скорости выполнения: о
//O(log n ), или логарифмическое время. Пример: бинарный поиск.
//о О(n), или линейное время. Пример: простой поиск.
//Q
//О(n * log n). Пример: эффективные алгоритмы сортировки (быстрая
//сортировка - но об этом в главе 4).
//о О(n 2 ). Пример: медленные алгоритмы сортировки (сортировка выбором -
//см. главу 2).
//о О(n!). Пример: очень медленные алгоритмы (задача о коммивояжере -
//о ней будет рассказано в следующем разделе).

func main() {
	//find k from sorted list, using binary search
	// Create a slice of 100 items
	//slice := make([]int, 10)
	//for i := 0; i < 10; i++ {
	//	slice[i] = i + rand.Intn(10)
	//}
	//
	//slices.Sort(slice)
	//fmt.Println(slice)
	s := generateSortedSlice(10)
	fmt.Println(s)
	search_binary(s, 4)
}

func search_binary(slice []int, target int) int {
	ilow := 0
	ihigh := len(slice) - 1

	fmt.Println(ilow, ihigh)
	var operations = 0

	if ilow == ihigh {
		if slice[ilow] == target {
			return target
		} else {
			return -1
		}
	}

	for ilow <= ihigh {
		imid := (ilow + ihigh) / 2
		if slice[imid] == target {
			fmt.Println("ilow:", ilow, "ihigh:", ihigh)
			fmt.Println("target number index", imid)
			fmt.Println("operations:", operations)
			return imid
		}
		if slice[imid] < target {
			ilow = imid + 1
			fmt.Println("ilow:", ilow, "ihigh:", ihigh)
			operations += 1
		} else {
			ihigh = imid - 1
			fmt.Println("ilow:", ilow, "ihigh:", ihigh)
			operations += 1
		}
	}
	fmt.Println("target number is not found")
	return operations
}

func generateSortedSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(100) // Generate random numbers between 0 and 999
	}
	sort.Ints(slice)
	return slice
}
