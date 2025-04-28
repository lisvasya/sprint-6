package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert определяет тип переданной строки и конвертирует её в соответствующий формат
func Convert(input string) (string, error) {
	// Убираем лишние пробелы по бокам строки
	input = strings.TrimSpace(input)

	// Если строка состоит только из символов Морзе (точки, тире и пробелы),
	// то мы предполагаем, что это код Морзе, и конвертируем его в обычный текст
	if isMorseCode(input) {
		return morse.ToText(input), nil
	}

	// Если это обычный текст, то конвертируем его в код Морзе
	return morse.ToMorse(input), nil
}

// Проверка, является ли строка кодом Морзе
func isMorseCode(input string) bool {
	// Строка в коде Морзе должна содержать только точку, тире и пробелы
	for _, ch := range input {
		if ch != '.' && ch != '-' && ch != ' ' {
			return false
		}
	}
	return true
}
