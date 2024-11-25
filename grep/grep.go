package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type grepFlags struct {
	LineNumber  bool // -n
	Filenames   bool // -l
	IgnoreCase  bool // -i
	InvertMatch bool // -v
	ExactMatch  bool // -x
}

func ExactMatch(line, pattern string) bool {
	return line == pattern
}

func IgnoreCase(line, pattern string) bool {
	return regexp.MustCompile(strings.ToLower(pattern)).MatchString(strings.ToLower(line))
}

func InvertMatch(matches bool) bool {
	return !matches
}

func Search(pattern string, flags, files []string) []string {
	grepFlags := grepFlags{}
	for _, flag := range flags {
		switch flag {
		case "-n":
			grepFlags.LineNumber = true
		case "-l":
			grepFlags.Filenames = true
		case "-i":
			grepFlags.IgnoreCase = true
		case "-v":
			grepFlags.InvertMatch = true
		case "-x":
			grepFlags.ExactMatch = true
		}
	}

	foundStrings := make([]string, 0)
	for _, file := range files {

		fileData, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file: ", err)
		}
		defer fileData.Close()

		var lines []string
		scanner := bufio.NewScanner(fileData)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())

		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning file: ", err)
		}
		multipleFiles := len(files) > 1

		for n, line := range lines {
			match := false

			if grepFlags.ExactMatch && grepFlags.IgnoreCase {
				match = strings.EqualFold(line, pattern)
			} else if grepFlags.ExactMatch {
				match = ExactMatch(line, pattern)
			} else if grepFlags.IgnoreCase {
				match = IgnoreCase(line, pattern)
			} else {
				match = strings.Contains(line, pattern)
			}

			if grepFlags.InvertMatch {
				match = InvertMatch(match)
			}

			if grepFlags.Filenames && match {
				foundStrings = append(foundStrings, file)
				break
			}

			if match {
				if grepFlags.LineNumber {
					if multipleFiles {
						foundStrings = append(foundStrings, fmt.Sprintf("%s:%d:%s", file, n+1, line))
					} else {
						foundStrings = append(foundStrings, fmt.Sprintf("%d:%s", n+1, line))
					}
				} else {
					if multipleFiles {
						foundStrings = append(foundStrings, fmt.Sprintf("%s:%s", file, line))
					} else {
						foundStrings = append(foundStrings, line)
					}
				}
			}
		}

	}
	return foundStrings
}
