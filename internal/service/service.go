package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	cleaned := strings.Trim(s, ".- ")
	return cleaned == ""
}

func Convert(data string) string {
	if isMorse(data) {
		return morse.ToText(data)
	}
	return morse.ToMorse(data)
}
