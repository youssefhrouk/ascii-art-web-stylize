package ascii

import "strings"

func HandleNewLines(splittedtxt []string, bannerPath string) string {
	var result strings.Builder

	for _, line := range splittedtxt {
		if line == "" {
			for i := 0; i <= 8; i++ {
				result.WriteString("\n")
			}
		} else {
			result.WriteString(Printer(Converter(line, FileReader(bannerPath))))
		}
	}

	return result.String()
}
