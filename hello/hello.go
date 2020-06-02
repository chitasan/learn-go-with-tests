package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const italian = "Italian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const italianHelloPrefix = "Ciao, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}


// private function as it starts with lowercase
func greetingPrefix(language string) (prefix string) {
    switch language {
    case spanish:
        prefix = spanishHelloPrefix
    case french:
        prefix = frenchHelloPrefix
    case italian:
        prefix = italianHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("world", ""))
}

// (prefix string) is a named return value
// named return value creates a variable in the function with a 'zero' value and return with 'return' 
// in Go, public functions start with capital letter, and private functions start with lowercase letter
