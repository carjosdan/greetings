package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	sayHello := fmt.Sprintf(randomHello(), name)
	return sayHello, nil
}

func ManyHello(names []string) (map[string]string, error) {
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

func randomHello() string {
	formats := []string{
		"Hello, Welcome %s! \n",
		"Hi %s, nice to see you! \n",
		"Greetings %s! \n",
		"Salutations %s! \n",
	}

	return formats[rand.Intn(len(formats))]
}
