package util

import (
	"strings"
	"unicode"
)

func addColonPerhaps(name string) string {
	if strings.HasPrefix(name, ":") {
		return name
	}
	return ":" + name
}

func removeColonPerhaps(name string) string {
	if strings.HasPrefix(name, ":") {
		return name[1:]
	}
	return name
}

func UrlToCanonical(name string) string {
	return removeColonPerhaps(
		strings.ToLower(strings.ReplaceAll(name, " ", "_")))
}

func DisplayToCanonical(name string) string {
	return removeColonPerhaps(
		strings.ToLower(strings.ReplaceAll(name, " ", "_")))
}

func CanonicalToDisplay(name string) (res string) {
	tmp := strings.Title(name)
	var afterPoint bool
	for _, ch := range tmp {
		if afterPoint {
			afterPoint = false
			ch = unicode.ToLower(ch)
		}
		switch ch {
		case '.':
			afterPoint = true
		case '_':
			ch = ' '
		}
		res += string(ch)
	}
	return addColonPerhaps(res)
}
