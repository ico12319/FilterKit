package main

import (
	"filters/craftIntersectionFilter"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	filter4 := craftIntersectionFilter.CreateIntersectionFilter("allowed.txt", "blocked.txt", 6)
	fmt.Println(filter4.Accepts("gopher"))
	fmt.Println(time.Since(startTime)) //for measuring performance
}
