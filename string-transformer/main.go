package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var smallWords = map[string]bool{
	"a": true, "an": true, "the": true, "and": true, "but": true,
	"or": true, "for": true, "nor": true, "on": true, "at": true,
	"to": true, "by": true, "in": true, "of": true, "up": true,
	"as": true, "is": true, "it": true,
}

var commands = map[string]func(string) string{}

func main() {
	commands = map[string]func(string) string{
		"upper":   strings.ToUpper,
		"lower":   strings.ToLower,
		"cap":     capWords,
		"title":   titleWords,
		"snake":   snakeCase,
		"reverse": reverseWords,
	}

	validCmds := []string{"upper", "lower", "cap", "title", "snake", "reverse", "exit"}

	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("--- ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := strings.ToLower(parts[0])
		text := strings.Join(parts[1:], " ")

		if cmd == "exit" {
			fmt.Println("Goodbye my dear!")
			return
		}

		if fn, ok := commands[cmd]; ok {
			if strings.TrimSpace(text) == "" {
				fmt.Printf("✗ No text provided. Usage: %s <text>\n", cmd)
				continue
			}
			fmt.Println("  ", fn(text))
		} else {
			fmt.Printf("✗ Unknown command: \"%s\"\n", cmd)
			fmt.Printf("  Valid commands: %s\n", strings.Join(validCmds, ", "))
		}
	}
}

func capWords(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
	}
	return strings.Join(words, " ")
}

func titleWords(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		l := strings.ToLower(w)
		if i == 0 || !smallWords[l] {
			words[i] = strings.ToUpper(string(l[0])) + l[1:]
		} else {
			words[i] = l
		}
	}
	return strings.Join(words, " ")
}

func snakeCase(text string) string {
	var b strings.Builder
	for _, r := range text {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '_' {
			b.WriteRune(toLower(r))
		} else if r == ' ' {
			b.WriteRune('_')
		}
	}
	return b.String()
}

func reverseWords(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		r := []rune(w)
		for j, k := 0, len(r)-1; j < k; j, k = j+1, k-1 {
			r[j], r[k] = r[k], r[j]
		}
		words[i] = string(r)
	}
	return strings.Join(words, " ")
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}
