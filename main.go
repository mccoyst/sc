package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var line []byte
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if line == nil {
			line = make([]byte, len(s.Bytes()))
			copy(line, s.Bytes())
		} else {
			line = append(line, s.Bytes()...)
		}

		if len(line) == 0 {
			line = nil
			continue
		}

		if line[len(line)-1] == '\\' {
			line = line[:len(line)-1]
			continue
		}

		parts := strings.Fields(string(line))
		c := exec.Command(parts[0], parts[1:]...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "! error when running %q: %v\n", line, err)
			break
		}
		line = nil

	}

	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "! error while reading input: %v\n", err)
		os.Exit(1)
	}
}
