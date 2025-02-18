// Exercice from: https://go.dev/tour/moretypes/23

package main

import "fmt"

func main() {
	line1 := "I am learning Go"
	map1 := map[string]int{"Go": 1, "I": 1, "am": 1, "learning": 1}

	line2 := "The quick brown fox jumped over the lazy dog"
	map2 := map[string]int{"The": 1, "brown": 1, "dog": 1, "fox": 1, "jumped": 1, "lazy": 1, "over": 1, "quick": 1, "the": 1}

	line3 := "To go to the mountain I need to go to another country"
	map3 := map[string]int{"I": 1, "To": 1, "another": 1, "country": 1, "go": 2, "mountain": 1, "need": 1, "the": 1, "to": 3}

	if sameMap(map1, wordCount(line1)) && sameMap(map2, wordCount(line2)) && sameMap(map3, wordCount(line3)) {
		fmt.Println("All good.")
	}
}

func sameMap(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, val := range map1 {
		if map2[key] != val {
			return false
		}
	}
	return true
}

func wordCount(s string) map[string]int {
	// TODO: return a map of the counts of each “word” in the string s (Case sensitive)
	// Hint: strings.Fields might be useful: https://pkg.go.dev/strings#Fields
}
