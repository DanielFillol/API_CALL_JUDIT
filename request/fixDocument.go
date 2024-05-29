package request

import (
	"strings"
)

func fixDocument(document string) string {
	if len(document) >= 11 {
		// If the length is already 11 or more, no need to modify.
		return document
	}

	// Calculate the number of zeros to add.
	numZeros := 11 - len(document)

	// Create a string with the required number of leading zeros.
	paddedDocument := strings.Repeat("0", numZeros) + document

	return paddedDocument
}
