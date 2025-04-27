package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// функция принимает строку, которая может быть текстом или кодом, определяет тип и переконвертирует в ддругой формат
func Convert(input string) (string, error) {
	// символы для характерные для морзе
	morseChars := ".- "
	// проверка состоит ли строка только из символов морзе
	isMorse := func(s string) bool {
		for _, r := range s {
			if !strings.ContainsRune(morseChars, r) {
				return false
			}
		}
		return true
	}

	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("empty input")
	}

	if isMorse(trimmed) {
		text := morse.ToText(trimmed)
		if text == "" {
			return "", errors.New("failed to convert morse to text")
		}
		return text, nil
	}

	morseCode := morse.ToMorse(trimmed)
	if morseCode == "" {
		return "", errors.New("failed to convert text to morse")
	}
	return morseCode, nil
}
