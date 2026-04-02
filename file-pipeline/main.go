package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] == os.Args[2] {
		fmt.Println("Usage: go run . input.txt output.txt> (Files must differ)")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("✗ Error reading %s\n", os.Args[1])
		return
	}

	lines := strings.Split(string(data), "\n")
	var out []string
	out = append(out, "SENTINEL FIELD REPORT — PROCESSED")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.Trim(line, "-") == "" {
			continue
		}

		line = strings.ReplaceAll(line, "TODO:", "✦ ACTION:")
		line = strings.ReplaceAll(line, "CLASSIFIED:", "[REDACTED]:")

		if strings.Contains(line, "REVERSE") {
			w := strings.Fields(line)
			for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
				w[i], w[j] = w[j], w[i]
			}
			line = strings.Join(w, " ")
		}

		out = append(out, fmt.Sprintf("%03d. %s", len(out), line))
	}

	if err := os.WriteFile(os.Args[2], []byte(strings.Join(out, "\n")+"\n"), 0644); err != nil {
		fmt.Println("✗ Error writing output.")
		return
	}
	fmt.Printf("✦ Processed %d lines into %s\n", len(out)-1, os.Args[2])
}
