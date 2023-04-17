package main

import "strings"

package main

// не удаляйте импорты, они используются при проверке
import (
"fmt"
"math/rand"
"os"
"strings"
"testing"
)

type Words struct {
	str string
	words []string
}

func MakeWords(s string) Words {
	return Words{s, strings.Fields(s)}
}

func (w Words) Index(word string) int {
	for idx, item := range w.words {
		if item == word {
			return idx
		}
	}
	return -1
}