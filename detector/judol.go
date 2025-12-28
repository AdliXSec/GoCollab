package detector

import (
	"regexp"
	"strings"
)

var judolKeywords = []string{
	"slot", "togel", "bet", "casino", "gacor",
	"maxwin", "judi", "taruhan",
}

var suspiciousTLD = []string{
	".xyz", ".top", ".click", ".site", ".live",
}

func IsJudol(url string) bool {
	url = strings.ToLower(url)

	// 1. Keyword detection
	for _, keyword := range judolKeywords {
		if strings.Contains(url, keyword) {
			return true
		}
	}

	// 2. Regex pattern detection
	regex := regexp.MustCompile(`(slot|bet|casino)[0-9]*`)
	if regex.MatchString(url) {
		return true
	}

	// 3. Suspicious TLD detection
	for _, tld := range suspiciousTLD {
		if strings.HasSuffix(url, tld) {
			return true
		}
	}

	return false
}
