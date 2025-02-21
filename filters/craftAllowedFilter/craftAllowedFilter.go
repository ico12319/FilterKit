package craftAllowedFilter

import (
	"bufio"
	"filters/allowed"
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

func CreateAllowedFilter(fileName string) filter.Filter {
	allowedStrings := readStringsFromFile(fileName)
	allowedFilter := allowed.NewAllowedInstance(allowedStrings)
	toReturn := filterByCriteria.NewFilterByCriteria(allowedFilter)
	return toReturn

}
