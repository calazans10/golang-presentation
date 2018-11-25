package hello

const french = "French"
const portuguese = "Portuguese"
const spanish = "Spanish"

// const italian = "Italian"

const defaultPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "

// const italianHelloPrefix = "Ciao, "

// Greeting returns a greeting message
func Greeting(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case french:
		return frenchHelloPrefix
	case portuguese:
		return portugueseHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return defaultPrefix
	}
}
