package ascii

import (
	"strings"
)

func PrintAndSplit(input, banner string) string {
	if len(input) == 0 {
		return "No input provided"
	}

	validArg := ValidateInput(input)
	if validArg == "" {
		return "Invalid input"
	}

	splited := strings.Split(validArg, string([]byte{13, 10}))
	bannerPath := GetBannerPath(banner)
	if strings.HasPrefix(bannerPath, "Error:") {
		return bannerPath
	}

	return HandleNewLines(splited, bannerPath)
}
