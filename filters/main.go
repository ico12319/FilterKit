package main

import (
	"filters/craftIntersectionFilter"
	"fmt"
)

func main() {
	filter4 := craftIntersectionFilter.CreateIntersectionFilter("allowed.txt", "blocked.txt", 6)
	fmt.Println(filter4.Accepts("gopher"))

}
