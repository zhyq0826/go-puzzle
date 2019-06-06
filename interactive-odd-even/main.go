package main

import (
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Getenv("n"))
}
