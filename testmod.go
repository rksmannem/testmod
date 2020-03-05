package testmod

import (
	"errors"
	"fmt"
)

// Hi ....
func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!\n", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!\n", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s!\n", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!\n", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
