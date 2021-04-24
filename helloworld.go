package main

// this is an import block, it allows multiple packages to be imported without
// repeating the "import" keyword
import (
	"fmt"
	"strings"
)

/*
The main function, when part of the main package, identifies the entry point of
an application.
*/
func main() {
	fmt.Println("Hello, playground")

	arrayPlay()
}

func arrayPlay() {

	var q [3]int = [3]int{1, 2, 3}

	fmt.Printf("The values of the q array\n")
	for _, v := range q {
		fmt.Printf("%d\n", v)
	}

	// Test for shifting an array
	s := []int{0, 1, 2, 3, 4, 5}
	// shift left by two positions
	fmt.Printf("Initial values of s, before shifting: \n%v\n", s)
	reverse(s[:2])
	fmt.Printf("After the first reverse: %v\n", s)
	reverse(s[2:])
	fmt.Printf("After the second reverse: %v\n", s)
	reverse(s)
	fmt.Printf("Final value: %v\n", s)
	fmt.Printf("Spinwords is: " + SpinWords("Test multi string"))

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SpinWords(str string) string {
	words := strings.Split(str, " ")

	newWords := make([]string, len(words))

	for i, r := range words {
		if len(r) >= 5 {
			newWords[i] = reverseWord(r)
		} else {
			newWords[i] = r
		}
	}
	return strings.Join(newWords, " ")

} // SpinWords

func reverseWord(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
