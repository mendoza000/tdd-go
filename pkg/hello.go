package pkg

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	switch lang {

	case spanish:
		if name == "" {
			name = "Mundo"
		}
		return spanishHelloPrefix + name + "!"

	case french:
		if name == "" {
			name = "Monde"
		}
		return frenchHelloPrefix + name + "!"

	default:
		if name == "" {
			name = "World"
		}
		return englishHelloPrefix + name + "!"
	}
}
