package main

// A concrete implementation of languageBot for the Spanish Language
type spanishBot struct{}

func (spanishBot) getGoodbye() string {
	return "Ciao"
}

func (spanishBot) getGreeting() string {
	return "Ola Mundo!!!"
}
