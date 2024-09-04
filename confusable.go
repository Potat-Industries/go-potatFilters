package potatFilters

import "strings"

// Replaces characters that might be easily confused with their English equivalents.
// Also removes accents (Unicode range U+0300 to U+036F) and standardizes most whitespace characters to \s.
func ReplaceConfusable(from string) string {
	var cleaned strings.Builder

	for _, v := range from {
		replaceWith, replace := replacements[v]

		if !replace {
			cleaned.WriteRune(v)

			continue
		}

		if replaceWith == "" {
			continue
		}

		cleaned.WriteString(replaceWith)
	}

	return cleaned.String()
}
