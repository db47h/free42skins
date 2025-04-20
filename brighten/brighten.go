package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func l(c int) int {
	return int(float64(c) / 0.85)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cs := bytes.Split(scanner.Bytes(), []byte{' '})
		r, err := strconv.Atoi(string(cs[0]))
		if err != nil {
			panic(err)
		}
		g, err := strconv.Atoi(string(cs[1]))
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(string(cs[2]))
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %d %d\n", l(r), l(g), l(b))
	}
}
