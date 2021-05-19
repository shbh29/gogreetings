package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v, Wecome!", name)
	return message
}
