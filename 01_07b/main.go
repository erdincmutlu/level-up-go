package main

import (
	"flag"
	"log"
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	brackets := ""
	for _, c := range expr {
		if c == '(' || c == '[' || c == '{' {
			brackets += string(c)
			continue
		}

		if (c == ')' && brackets[len(brackets)-1] == '(') ||
			(c == ']' && brackets[len(brackets)-1] == '[') ||
			(c == '}' && brackets[len(brackets)-1] == '{') {
			brackets = brackets[:len(brackets)-1]
			continue
		}

		if c == ')' || c == ']' || c == '}' {
			return false
		}
	}

	return len(brackets) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
