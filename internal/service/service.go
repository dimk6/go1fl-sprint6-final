package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("empty file")
	}

	if isMorse(data) {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}

func isMorse(data string) bool {
	for _, r := range data {
		if r != '.' && r != '-' && r != ' ' && r != '\n' && r != '\r' && r != '\t' {
			return false
		}
	}
	return true
}
