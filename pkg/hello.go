package pkg

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, lang string) string {
	if lang == spanish {
		if name == "" {
			name = "Mundo"
		}
		return spanishHelloPrefix + name + "!"
	}

	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + "!"
}
