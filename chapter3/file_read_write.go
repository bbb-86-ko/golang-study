package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	example_write_to()
	fmt.Println()
	example_read_from()
	fmt.Println()
}

func file_read_example_1() error {
	f, err := os.Open("file_read.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	var num int
	if _, err := fmt.Fscanf(f, "%d", &num); err == nil {

	}
	return nil
}

func file_write_example_1() error {
	f, err := os.Create("file_write.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	var num int
	if _, err := fmt.Fscanf(f, "%d", &num); err != nil {
		return err
	}
	return nil
}

func file_write_to(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func file_read_from(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func example_write_to() {
	lines := []string{
		"xxx@email.com",
		"yyy@email.com",
		"zzz@email.com",
	}

	if err := file_write_to(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
}

func example_read_from() {
	r := strings.NewReader("file_read.txt")
	var lines []string
	if err := file_read_from(r, &lines); err != nil {
		fmt.Println(err)
	}

	fmt.Println(lines)
}
