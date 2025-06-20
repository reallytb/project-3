package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) string {
	morseChars := ".- "
	isText := true
	for _, char := range data {
		if strings.ContainsRune(morseChars, char) {
			isText = false
		}
	}
	switch isText {
	case true:
		morse := morse.ToMorse(data)
		return morse
	case false:
		text := morse.ToText(data)
		return text
	}
	log.Fatal("ошибка конвертации текста/морзе")
	fmt.Println("ошибка конвертации текста/морзе")
	return ""
}
