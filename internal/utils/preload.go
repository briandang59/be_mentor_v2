package utils

import (
	"net/url"
	"strings"
	"unicode"
)

// chuyển "cover-letters" -> "CoverLetters"
func kebabToPascal(input string) string {
	parts := strings.Split(input, "-")
	for i, part := range parts {
		if len(part) > 0 {
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			parts[i] = string(runes)
		}
	}
	return strings.Join(parts, "")
}

// chuyển "tags.cover-letters" -> "Tags.CoverLetters"
func capitalizeEachPart(path string) string {
	parts := strings.Split(path, ".")
	for i, part := range parts {
		parts[i] = kebabToPascal(part)
	}
	return strings.Join(parts, ".")
}

// ParsePopulateQuery nhận URL params và trả về list các field cần preload
// Ví dụ:
// ?populate=tags&populate=cover-letters
// ?populate[tags]=cover-letters
// => []string{"Tags", "CoverLetters", "Tags.CoverLetters"}
func ParsePopulateQuery(values url.Values) []string {
	preloadMap := make(map[string]bool)

	// dạng populate[tags][cover-letters]
	for key := range values {
		if strings.HasPrefix(key, "populate[") {
			str := strings.TrimPrefix(key, "populate[")
			str = strings.TrimSuffix(str, "]")
			str = strings.ReplaceAll(str, "][", ".") // "tags][cover-letters" -> "tags.cover-letters"

			normalized := capitalizeEachPart(str)
			preloadMap[normalized] = true
		}
	}

	// dạng populate=tags&populate=cover-letters
	for _, populateValue := range values["populate"] {
		normalized := capitalizeEachPart(populateValue)
		preloadMap[normalized] = true
	}

	// gom kết quả
	var result []string
	for key := range preloadMap {
		result = append(result, key)
	}
	return result
}
