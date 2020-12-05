package main

import "fmt"

// DictionaryNumber is...
type DictionaryNumber struct {
	numbers [10]Number
}

// Number is ...
type Number struct {
	name    string
	display [5]string
}

// Saludar is ...
func (p Number) Saludar() {
	fmt.Printf("Hola, mi nombre es %s\n%s\n%s\n%s\n%s\n%s\n", p.name, p.display[0], p.display[1], p.display[2], p.display[3], p.display[4])
}

//ConcatNumber is ...
func ConcatNumber(d DictionaryNumber, s []int) {
	fmt.Printf("Numero: %v , Tamaño: %d\n", s, len(s))
	auxRow := [5]string{"", "", "", "", ""}
	n := 0
	for i := 0; i < 5; i++ {
		for p := 0; p < len(s); p++ {
			n = s[p]
			auxRow[i] = auxRow[i] + d.numbers[n].display[i] + ""
		}
	}
	fmt.Printf("%s\n%s\n%s\n%s\n%s\n", auxRow[0], auxRow[1], auxRow[2], auxRow[3], auxRow[4])
}

func main() {
	var dict DictionaryNumber
	dict.numbers[0].name = "numero 0"
	dict.numbers[0].display = [5]string{" - - - - ", " -     - ", " -     - ", " -     - ", " - - - - "}
	dict.numbers[1].name = "numero 1"
	dict.numbers[1].display = [5]string{"       - ", "       - ", "       - ", "       - ", "       - "}
	dict.numbers[2].name = "numero 2"
	dict.numbers[2].display = [5]string{" - - - - ", "       - ", " - - - - ", " -       ", " - - - - "}
	dict.numbers[3].name = "numero 3"
	dict.numbers[3].display = [5]string{" - - - - ", "       - ", " - - - - ", "       - ", " - - - - "}
	dict.numbers[4].name = "numero 4"
	dict.numbers[4].display = [5]string{" -     - ", " -     - ", " - - - - ", "       - ", "       - "}
	dict.numbers[5].name = "numero 5"
	dict.numbers[5].display = [5]string{" - - - - ", " -       ", " - - - - ", "       - ", " - - - - "}
	dict.numbers[6].name = "numero 6"
	dict.numbers[6].display = [5]string{" - - - - ", " -       ", " - - - - ", " -     - ", " - - - - "}
	dict.numbers[7].name = "numero 7"
	dict.numbers[7].display = [5]string{" - - - - ", "       - ", "       - ", "       - ", "       - "}
	dict.numbers[8].name = "numero 8"
	dict.numbers[8].display = [5]string{" - - - - ", " -     - ", " - - - - ", " -     - ", " - - - - "}
	dict.numbers[9].name = "numero 9"
	dict.numbers[9].display = [5]string{" - - - - ", " -     - ", " - - - - ", "       - ", "       - "}
	dict.numbers[1].Saludar()
	var s []int
	s = append(s, 4, 6, 5, 7, 8, 3, 5, 9)
	ConcatNumber(dict, s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
