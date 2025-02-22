# FilterKit
 FilterKit is a modular, composable system that applies multiple validation criteria to strings in Go.

# Filtering System

This repository provides a modular, composable filtering system in Go. The main goal is to determine whether a given string satisfies one or more validation criteria, using a clean, extensible design. Here's an overview of the key components:

## Overview

- **Filter Interface**  
  Each filter implements the `filter.Filter` interface, which declares a single method:
  
      type Filter interface {
          Accepts(s string) bool
      }
  
  This design allows different criteria to be wrapped as separate filters.

- **Allowed & Blocked Filters**  
  - `allowed.Allowed` loads a list of allowed strings and checks if an input exists in that list.  
  - `blocked.Blocked` does the opposite: it loads a list of blocked strings and returns `false` if the input is in that list.

- **Length Filter**  
  - `lengthBlocker.NewLengthBlocker(maxLength)` validates whether a string’s length is below a specified maximum.

- **Intersection of Filters**  
  - `intersectionOfFilters.IntersectionOfFilters` combines multiple filters and returns `true` only if _all_ included filters accept the input.

- **Crafting Filters**  
  - The “craft” packages (e.g., `craftAllowedFilter`, `craftBlockedFilter`, `craftMaxLengthFilter`) read external files (each line treated as a string entry) and create corresponding filters.  
  - `craftIntersectionFilter` creates the final composite filter by adding an allowed filter, a blocked filter, and the length filter to `IntersectionOfFilters`.

## Usage

1. **Create text files** listing the allowed and blocked strings (one item per line).  
2. **Instantiate the composite filter** in your Go code:

       import (
           "filters/craftIntersectionFilter"
       )

       func main() {
           // Example usage:
           compositeFilter := craftIntersectionFilter.CreateIntersectionFilter(
               "allowed.txt",  // Path to file with allowed strings
               "blocked.txt",  // Path to file with blocked strings
               30,             // Max length
           )

           // Test any given string:
           testString := "SomeStringToTest"
           if compositeFilter.Accepts(testString) {
               // String passes all checks
           } else {
               // String fails one or more checks
           }
       }

3. **Extend or modify filters** to fit your needs—just implement `filter.Filter` for new validation criteria and add it to your intersection filter.

## Why This Design?

- **Modularity**: Each filter has a single responsibility, making the system easy to extend.  
- **Composability**: `IntersectionOfFilters` can combine any number of filters, enabling complex validation rules from simpler building blocks.  
- **Simplicity**: Each package handles a specific part of the filtering logic (loading data, validating strings, etc.), ensuring the code is straightforward to maintain.


## Running Tests To run the unit tests for FilterKit, follow these steps: 
1. **Open your terminal.** 
2. **Navigate to the project root directory.**: For example, if your project is in `/Users/yourname/FilterKit`, run: bash cd /Users/yourname/FilterKit 
3. **Run the tests using the Go test command:**: bash go test ./...  This command will search for all files ending with `_test.go` in your project and execute the tests.
 4. **Review the test output.** If all tests pass, you'll see output similar to: PASS ok filters/allowed 0.417s

