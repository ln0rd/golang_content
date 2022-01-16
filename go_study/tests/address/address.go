package address

import (
	"strings"
)

// kinf of address
func KindOfAddress(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "Rodovia"}
	firstWord := strings.Split(address, " ")[0]
	hasValidAddress := false

	for _, validType := range validTypes {
		if validType == firstWord {
			hasValidAddress = true
		}
	}

	if hasValidAddress {
		return strings.Title(firstWord)
	}

	return "invalid type"
}