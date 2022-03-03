package greetings

import (
	"errors"
	"fmt"
)

// Hello reutrns a greeting for the named person.
func Hello(name string) (string, error) { // Function name, Parameter type, Return type
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name) // Shortcut for declaring and initializing a variable
	return message, nil
}
