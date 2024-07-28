package ascii

import (
	"net/http"
	"strings"
)

func PrintAndSplit(input, banner string) (string, int) {
	validArg := ValidateInput(input)
	if validArg == "" {
		return "400 : Bad Request - Invalid input", http.StatusBadRequest
	}
	splited := strings.Split(validArg, "\r\n")
	bannerPath, isValid := GetBannerPath(banner)
	if !isValid {
		return "400 : Bad Request - Invalid banner", http.StatusBadRequest
	}
	return HandleNewLines(splited, bannerPath), http.StatusOK
}
