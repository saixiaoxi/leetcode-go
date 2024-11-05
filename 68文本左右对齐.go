package main

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var currentLine []string
	currentLength := 0

	for _, word := range words {
		if currentLength+len(word)+len(currentLine) > maxWidth {
			totalSpaces := maxWidth - currentLength
			if len(currentLine) == 1 {
				result = append(result, currentLine[0]+strings.Repeat(" ", totalSpaces))
			} else {
				spaceBetweenWords := totalSpaces / (len(currentLine) - 1)
				extraSpaces := totalSpaces % (len(currentLine) - 1)
				line := ""
				for i := 0; i < len(currentLine)-1; i++ {
					line += currentLine[i] + strings.Repeat(" ", spaceBetweenWords)
					if i < extraSpaces {
						line += " "
					}
				}
				line += currentLine[len(currentLine)-1]
				result = append(result, line)
			}
			currentLine = []string{}
			currentLength = 0
		}
		currentLine = append(currentLine, word)
		currentLength += len(word)
	}

	lastLine := strings.Join(currentLine, " ")
	lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
	result = append(result, lastLine)

	return result
}
