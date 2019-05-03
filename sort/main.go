package main

import "fmt"

func merge(T []int, ind int) {
	G := T[:ind]
	D := T[ind:]

	l := len(T)
	Q := make([]int, l)

	for k := 0; k < l; k++ {
		if len(G) == 0 {
			Q[k] = D[0]
			D = D[1:]
		} else if len(D) == 0 {
			Q[k] = G[0]
			G = G[1:]
		} else if G[0] <= D[0] {
			Q[k] = G[0]
			G = G[1:]
		} else {
			Q[k] = D[0]
			D = D[1:]
		}
	}
	for k := 0; k < l; k++ {
		T[k] = Q[k]
	}
}

func sort(T []int) {
	l := len(T)
	if l <= 1 {
		return
	}
	if l == 2 {
		if T[0] > T[1] {
			T[0], T[1] = T[1], T[0]
		}
		return
	}
	i := int(l / 2)
	sort(T[:i])
	sort(T[i:])
	merge(T, i)
}

func main() {
	T := []int{6, 4, 7, 2, 3, 9, 1, 5}
	//T := []int{4, 7, 2}

	fmt.Println(T)
	sort(T)
	fmt.Println(T)

}
