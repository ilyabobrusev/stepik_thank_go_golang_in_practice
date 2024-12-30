package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	res := text
	if len(text) > width {
		res = text[:width] + "..."
	}

	fmt.Println(res)
}
