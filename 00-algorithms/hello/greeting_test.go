package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		actual   string
	}{
		{"saying hello to people", "Hello, Jason", Greeting("Jason", "")},
		{"saying hello world when an empty string is supplied", "Hello, World", Greeting("", "")},
		{"saying hello to people in Spanish", "Hola, Julieta", Greeting("Julieta", "Spanish")},
		{"saying hello to people in French", "Bonjour, Brigitte", Greeting("Brigitte", "French")},
		{"saying hello to people in Portuguese", "Olá, Lucas", Greeting("Lucas", "Portuguese")},
		// {"saying hello to people in Italian", "Ciao, Buffon", Greeting("Buffon", "Italian")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.actual)
		})
	}
}
