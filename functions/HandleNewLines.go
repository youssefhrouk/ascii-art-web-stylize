package ascii

import "strings"

func HandleNewLines(splittedtxt []string, bannerPath string) string {
	var result strings.Builder
	empty := 0

	for _, line := range splittedtxt {
		if line == "" {
			empty++
		} else {
			if empty > 0 {
				for i := 0; i < empty; i++ {
					result.WriteString("\n")
				}
				empty = 0
			}
			result.WriteString(Printer(Converter(line, FileReader(bannerPath))))
		}
	}

	// Print trailing newlines if any
	if empty > 0 {
		for i := 0; i < empty; i++ {
			if i == len(splittedtxt)-1 {
				break
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
