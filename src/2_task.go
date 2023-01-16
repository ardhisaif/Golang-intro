package src

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

func GeneratePassword(password string, level string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	var result string
	var medium = []rune("~`!@#$%^&*()_+-=")
	var strong = []rune("~`!@#$%^&*()_+-=1234567890")
	if len(password) < 6 {
		return "", errors.New("panjang password harus lebih dari 6")
	}

	for i := 0; i < len(password); i++ {
		if i != 0 {
			if rand.Intn(len(password))%2 == 0 {
				result += strings.ToLower(string(password[i]))
			} else {
				result += strings.ToUpper(string(password[i]))
			}
		} else {
			result += strings.ToUpper(string(password[i]))
		}
	}

	if level == "low" {
		return string(result), nil
	}

	if level == "medium" {
		b := make([]rune, 5)
		for i := range b {
			b[i] = medium[rand.Intn(len(medium))]
		}
		result = string(result) + string(b)
	}

	if level == "strong" {
		b := make([]rune, 10)
		for i := range b {
			b[i] = strong[rand.Intn(len(strong))]
		}
		result = string(result) + string(b)
	}

	return result, nil
}
