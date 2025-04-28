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
	// Проверим, является ли строка кодом Морзе
	if isMorse(input) {
		// Если да, конвертируем в текст
		return morse.ToText(input), nil
	} else if isText(input) {
		// Если это текст, конвертируем в код Морзе
		return morse.ToMorse(input), nil
	} else {
		return "", errors.New("failed to determine the data format")
	}
}

// Функция для проверки, является ли строка кодом Морзе
func isMorse(input string) bool {
	// В коде Морзе могут быть только точки, тире и пробелы
	// Пример: .- .-.. .-.. ---
	// Проверим, есть ли в строке только символы '.', '-', ' ' (пробел)
	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' {
			return false
		}
	}
	return true
}

// Функция для проверки, является ли строка обычным текстом
func isText(input string) bool {
	// Если строка состоит из обычных символов (буквы, пробелы, цифры), то это текст
	return strings.TrimSpace(input) != ""
}
