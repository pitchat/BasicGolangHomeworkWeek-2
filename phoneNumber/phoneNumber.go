package main

import (
	"fmt"
	"strings"
)

func phonebook(phoneNoList []string) map[string]int {
	phoneBook := make(map[string]int)
	for _, phoneNo := range phoneNoList {
		phoneNo = strings.ReplaceAll(phoneNo, "(", "")
		phoneNo = strings.ReplaceAll(phoneNo, ")", "")
		phoneNo = strings.ReplaceAll(phoneNo, "-", "")
		phoneNo = strings.ReplaceAll(phoneNo, " ", "")
		phoneBook[phoneNo]++
	}
	return phoneBook
}

func main() {

	phoneNoList := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892"}
	phoneBook := phonebook(phoneNoList)
	fmt.Printf("%#v\n", phoneBook)

}
