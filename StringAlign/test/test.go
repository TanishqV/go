package main

import (
	"fmt"
	"StringAlign/stralign"
)

func main() {
	var sample = []string{"some", " things", "to ", " test ", "Something larger than width"}

	for _,s := range sample {
		fmt.Printf("'%s'\n", s)
		fmt.Printf("Left justify:   '%s'\n", stralign.Ljust(s, 15, " "))
		fmt.Printf("Right justify:  '%s'\n", stralign.Rjust(s, 15, "-"))
		fmt.Printf("Center justify: '%s'\n", stralign.Center(s, 15, "_"))
		fmt.Printf("\n\n")
	}
}
