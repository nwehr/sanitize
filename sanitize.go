package sanitize

import (
	"fmt"
	"regexp"
)

func String(str string) (string, error) {
	re := regexp.MustCompile(`(?im)insert|update|drop|delete|truncate|add|create|replace|insert|commit|grant|constraint|set`)

	if re.MatchString(str) {
		return "", fmt.Errorf("string contains unsafe characters")
	}

	return str, nil
}
