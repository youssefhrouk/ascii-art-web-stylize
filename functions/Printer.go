package ascii

import "strings"

func Printer(ascLine [][]string) string {
	var result strings.Builder
	if len(ascLine) == 0 {
		return ""
	}
	for line := 0; line < len(ascLine[0]); line++ {
		for char := 0; char < len(ascLine); char++ {
			result.WriteString(ascLine[char][line])
		}
		result.WriteString("\n")
	}
	return result.String()
}
