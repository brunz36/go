package main

// A concrete implementation of languageBot for the English Language
type englishBot struct{}

func (englishBot) getGoodbye() string {
	return "Goodbye"
}

func (englishBot) getGreeting() string {
	return "Hello World!!!"
}
