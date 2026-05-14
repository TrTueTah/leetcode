func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	wordIndex := 0
	totalWords := len(words)

	// Process words until all are consumed
	for wordIndex < totalWords {

		// Collect words that fit in current line
		currentLineWords := []string{words[wordIndex]}
		currentLineLength := len(words[wordIndex])
		wordIndex++

		// Add more words while they fit
		for wordIndex < totalWords &&
			currentLineLength+1+len(words[wordIndex]) <= maxWidth {

			currentLineLength += 1 + len(words[wordIndex])
			currentLineWords = append(currentLineWords, words[wordIndex])
			wordIndex++
		}

		// Last line OR single word -> left justified
		if wordIndex == totalWords || len(currentLineWords) == 1 {

			line := strings.Join(currentLineWords, " ")

			padding := strings.Repeat(" ", maxWidth-len(line))

			result = append(result, line+padding)

			continue
		}

		// Calculate spaces
		totalChars := 0

		for _, word := range currentLineWords {
			totalChars += len(word)
		}

		totalSpacesNeeded := maxWidth - totalChars

		gaps := len(currentLineWords) - 1

		baseSpacesBetweenWords := totalSpacesNeeded / gaps

		extraSpaces := totalSpacesNeeded % gaps

		// Build justified line
		var currentLine strings.Builder

		for i := 0; i < len(currentLineWords)-1; i++ {

			currentLine.WriteString(currentLineWords[i])

			spacesToAdd := baseSpacesBetweenWords

			if i < extraSpaces {
				spacesToAdd++
			}

			currentLine.WriteString(strings.Repeat(" ", spacesToAdd))
		}

		// Add last word
		currentLine.WriteString(currentLineWords[len(currentLineWords)-1])

		result = append(result, currentLine.String())
	}

	return result
}