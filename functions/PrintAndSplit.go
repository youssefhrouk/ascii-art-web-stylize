package ascii

import (
	"strings"
)

func PrintAndSplit(input, banner string) string {
	validArg := ValidateInput(input)
	if validArg == "" {
		return "Invalid input"
	}
	splited := strings.Split(validArg, "\r\n")
	bannerPath := GetBannerPath(banner)
	return HandleNewLines(splited, bannerPath)
}
