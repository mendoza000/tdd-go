package pkg

const french = "French"
const spanish = "Spanish"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

func gettingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func gettingWorld(lang string) (world string) {
	switch lang {
	case spanish:
		world = "Mundo"
	case french:
		world = "Monde"
	case portuguese:
		world = "Mundo"
	default:
		world = "World"
	}
	return
}

func Hello(name string, lang string) string {
	if name == "" {
		name = gettingWorld(lang)
	}

	return gettingPrefix(lang) + name + "!"
}
