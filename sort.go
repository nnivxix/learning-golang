package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name  string
	Price int
}

type Products []Product

// Implement Sort Interface
func (p Products) Len() int {
	return len(p)
}

func (p Products) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Products) Less(i, j int) bool {
	// can change the condition (permanently)
	return p[i].Price > p[j].Price
}

func main() {
	var data []int = []int{4, 6, 2, 4, 1, 8, 5}

	fmt.Println("original data: ", data)
	sort.Ints(data)

	fmt.Println("ascending order: ", data)

	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	fmt.Println("descending order: ", data)

	products := Products{
		{"Sabun", 2000},
		{"Permen", 500},
		{"Kue", 1500},
	}

	sort.Sort(Products(products))
	fmt.Println(products)

}
