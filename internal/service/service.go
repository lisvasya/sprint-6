package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert определяет тип переданной строки и конвертирует её в соответствующий формат
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("empty line")
	}

	// Определяем, что это — Морзе или текст.
	// Если найдём символы, которые не входят в этот набор — считаем текстом.
	morseChars := ".- "
	isMorse := true
	for _, r := range input {
		if !strings.ContainsRune(morseChars, r) {
			isMorse = false
			break
		}
	}

	if isMorse {
		// Конвертируем из Морзе в текст
		text := morse.ToText(input)
		if strings.TrimSpace(text) == "" {
			return "", errors.New("incorrect Morze code")
		}
		return text, nil
	} else {
		// Конвертируем из текста в Морзе
		morseCode := morse.ToMorse(input)
		if strings.TrimSpace(morseCode) == "" {
			return "", errors.New("incorrect text for conversion")
		}
		return morseCode, nil
	}
}
