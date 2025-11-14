package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Julo", "Greg", "Asep", "Zola"}
	number := []int{89, 23, 20, 99}

	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(number))

	fmt.Println(names)
	slices.Sort(names)
	fmt.Println(names)

	// // Cara pakai slices.Values dengan iterator
	// fmt.Println("Iterating values:")
	// for v := range slices.Values(names) {
	// 	fmt.Println(v)
	// }

	// // Atau langsung loop tanpa Values (lebih simpel)
	// fmt.Println("\nDirect loop:")
	// for i, name := range names {
	// 	fmt.Println(name, i)
	// }

	// insert slice to names
	names = slices.Insert(names, len(names), "Hanasa")

	fmt.Println(names)

	// remove "Asep"
	asepIndex := slices.Index(names, "Asep")
	names = slices.Delete(names, asepIndex, 1)
	fmt.Println(names) // Asep deleted

}
