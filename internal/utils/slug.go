package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func GenerateSlug(input string) string {
	slug := strings.ToLower(input)

	slug = removeDiacritics(slug)

	slug = strings.ReplaceAll(slug, " ", "-")

	reg := regexp.MustCompile(`[^a-z0-9\-]+`)
	slug = reg.ReplaceAllString(slug, "")

	regDash := regexp.MustCompile(`-+`)
	slug = regDash.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}

func removeDiacritics(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case 'á', 'à', 'ả', 'ã', 'ạ', 'ă', 'ắ', 'ằ', 'ẳ', 'ẵ', 'ặ', 'â', 'ấ', 'ầ', 'ẩ', 'ẫ', 'ậ':
			b.WriteRune('a')
		case 'đ':
			b.WriteRune('d')
		case 'é', 'è', 'ẻ', 'ẽ', 'ẹ', 'ê', 'ế', 'ề', 'ể', 'ễ', 'ệ':
			b.WriteRune('e')
		case 'í', 'ì', 'ỉ', 'ĩ', 'ị':
			b.WriteRune('i')
		case 'ó', 'ò', 'ỏ', 'õ', 'ọ', 'ô', 'ố', 'ồ', 'ổ', 'ỗ', 'ộ', 'ơ', 'ớ', 'ờ', 'ở', 'ỡ', 'ợ':
			b.WriteRune('o')
		case 'ú', 'ù', 'ủ', 'ũ', 'ụ', 'ư', 'ứ', 'ừ', 'ử', 'ữ', 'ự':
			b.WriteRune('u')
		case 'ý', 'ỳ', 'ỷ', 'ỹ', 'ỵ':
			b.WriteRune('y')
		default:
			if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
				b.WriteRune(r)
			}
		}
	}
	return b.String()
}
