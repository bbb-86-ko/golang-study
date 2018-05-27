package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	example_readfrom_print()
}

func example_readfrom_1(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func example_readfrom_print() {
	r := strings.NewReader("bill tom jane")
	err := example_readfrom_1(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}

}
