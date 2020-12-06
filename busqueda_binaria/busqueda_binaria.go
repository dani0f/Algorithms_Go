package main

import (
	"fmt"
)

//BinarySearch is ...
func BinarySearch(s []int, element int) int {
	sizeArray := len(s)
	IndexL := 0
	IndexR := sizeArray - 1
	IndexM := (IndexR + IndexL) / 2
	found := 0
	for IndexR > IndexL {
		if s[IndexM] == element {
			found = 1
			break
		} else {
			if element < s[IndexM] {
				IndexR = IndexM - 1
			} else {
				IndexL = IndexM + 1
			}
		}
		IndexM = (IndexR + IndexL) / 2
	}
	if s[IndexM] == element && found == 0 {
		found = 1
	}

	if found == 1 {
		return IndexM
	}
	return -1

}

func main() {
	var s []int
	s = append(s, 1, 3, 4, 5, 6, 7, 8, 10)
	printSlice(s)
	element := 7
	response := BinarySearch(s, element)
	if response != -1 {
		fmt.Printf("Elemento %d encontrado en posición %d contando desde el 0, con valor de %d\n", element, response, s[response])
	} else {
		fmt.Print("No se encontro el elemento en el arreglo")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
