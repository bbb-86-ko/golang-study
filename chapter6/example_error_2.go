package main

func f() (int, error) {
	var err error
	return 1, err
}

func Must(i int, err error) int {
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	Must(f())
}
