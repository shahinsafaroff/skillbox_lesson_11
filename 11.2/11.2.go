package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "a10 10 20b 20 30c30 30 dd"
	regex := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	fmt.Printf("В строке содержатся числа в десятичном формате: %v\n", regex.MatchString(str))
	sub_mat_chall := regex.FindAllString(str, -1)
	for _, element := range sub_mat_chall {
		fmt.Println(element)
	}
}
