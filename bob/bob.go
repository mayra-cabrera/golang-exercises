package bob

import (
	"strings"
)

const testVersion = 2

func Hey(input string) string {
	normalize_input := strings.TrimSpace(input)
	switch {
	case strings.ToUpper(normalize_input) == normalize_input && strings.ToLower(normalize_input) != strings.ToUpper(normalize_input):
		return "Whoa, chill out!"
	case strings.HasSuffix(normalize_input, "?"):
		return "Sure."
	case normalize_input == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
