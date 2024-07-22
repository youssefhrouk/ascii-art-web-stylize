package ascii

import "strings"

func ValidateInput(arg string) string {
	var validChars strings.Builder
	for _, char := range arg {
		if char >= 32 && char <= 126 || char == 10 || char == 13 {
			validChars.WriteRune(char)
		}
	}
	return validChars.String()
}
