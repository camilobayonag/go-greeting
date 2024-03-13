package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello, devuevle un saludo para la persona espesifica
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Nombre Vacio")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"¡Holsa, %v! ¡Bienvenido!",
		"¡Que bueno verte, %v!",
		"¡Saludos, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}
