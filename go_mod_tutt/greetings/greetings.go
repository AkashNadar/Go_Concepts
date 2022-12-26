package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty Name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	format := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v",
		"Heil, %v well met!",
	}

	return format[rand.Intn(len(format))]
}
