package craftBlockedFilter

import (
	"bufio"
	blocked2 "filters/blocked"
	"filters/filter"
	"filters/filterByCriteria"
	"fmt"
	"os"
)

func readStringsFromFile(fileName string) map[string]int {
	extractedStrings := make(map[string]int)
	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error of type %s", err)
		return nil
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		token := scanner.Text()
		extractedStrings[token]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while scanning: %s\n", err)
		return nil
	}
	return extractedStrings
}

func CreateBlockedFilter(fileName string) filter.Filter {
	blockedStrings := readStringsFromFile(fileName)
	blocked := blocked2.NewBlockedInstance(blockedStrings)
	toReturn := filterByCriteria.NewFilterByCriteria(blocked)
	return toReturn
}
