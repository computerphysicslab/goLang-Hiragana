package main

import (
	"fmt"
)

func main() {
	// Learning goLang and 日本語 at the same time...
	// Unicode range:
	//	U+4E00–U+9FBF Kanji
	// 	U+3040–U+309F Hiragana
	// 	U+30A0–U+30FF Katakana
	firstRune := 'ぁ'-1
	numCols := 16
	numRows := 6
	for i := 0; i < numRows; i++ {
		fmt.Println("")
		for j := 0; j < numCols; j++ {
			fmt.Printf("%s ", string(firstRune +rune(j+i*numCols)))
		}
	}
}
