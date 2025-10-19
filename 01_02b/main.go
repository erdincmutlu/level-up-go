package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	for _, word := range words {
		var out []string
		for i, letter := range word {
			repeated := strings.Repeat(string(letter), i+1)
			out = append(out, repeated)
		}
		print(strings.Join(out, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
