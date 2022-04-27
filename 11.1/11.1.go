package main

import "fmt"

func main() {
	str := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	var upper = 0
	var i = 0
	for i = 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			upper++
		}
	}
	fmt.Printf("Количество слов с заглавной буквы: %d\n", upper)
}
