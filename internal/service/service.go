package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет, является ли входная строка текстом или кодом Морзе,
// и конвертирует её в другой формат.
// В противном случае считаем это обычным текстом и конвертируем туда.
// Convert пытается определить, является ли строка кодом Морзе или текстом.
// Возвращает строку, которая является результатом конвертации.
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("empty input")
	}

	// Проверяем, являются ли все символы допустимыми для кода Морзе
	if isMorse(input) {
		// Если это код Морзе, конвертируем в текст
		result := morse.ToText(input)
		if result == "" {
			// Если не удалось декодировать код Морзе, возвращаем ошибку
			return "", errors.New("failed to decode Morse code")
		}
		return result, nil
	}

	// Если это не код Морзе, конвертируем в код Морзе
	return morse.ToMorse(input), nil
}

// isMorse проверяет, является ли строка кодом Морзе
func isMorse(s string) bool {
	// Пробелы могут быть допустимы только для разделения символов в коде Морзе
	for _, r := range s {
		if !(r == '.' || r == '-' || r == ' ') {
			// Если встречаем символ, который не является точкой, дефисом или пробелом, это не код Морзе
			return false
		}
	}
	return true
}
