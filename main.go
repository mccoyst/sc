package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		parts := strings.Fields(line)
		c := exec.Command(parts[0], parts[1:]...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "! error when running %q: %v\n", line, err)
			break
		}
	}

	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "! error while reading input: %v\n", err)
		os.Exit(1)
	}
}
