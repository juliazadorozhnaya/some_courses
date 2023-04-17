package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	phrase := readString()
	words := strings.Fields(phrase)
	var w []string
	for _, word := range words {
		firstChar, _ := utf8.DecodeRuneInString(word)

		if unicode.IsLetter(firstChar) {
			w = append(w, string(firstChar))
		}
	}
	str := strings.Join(w, "")
	str = strings.ToUpper(str)
	fmt.Printf(str)
}

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
