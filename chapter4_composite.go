package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	values := []int{4, 7, 2, 5, 34, 2, 8, 9, 0}
	result := sort(values)
	for _, i := range result {
		fmt.Printf("%d\t", i)
	}
}

//binary tree sort

type tree struct {
	value       int
	left, right *tree
}

func sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	a := appendValues(values[:], root)
	return a

}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(r *tree, value int) *tree {
	if r == nil {
		r = new(tree)
		r.value = value
		return r
	}

	if value < r.value {
		r.left = add(r.left, value)
	} else {
		r.right = add(r.right, value)
	}

	return r

}

//Map
func count() {
	counts := make(map[rune]int)
	invalid := 0

	var utflen [utf8.UTFMax + 1]int
	in := bufio.NewReader(os.Stdin)
	for {

		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprint(os.Stderr, "character: %v\n", err)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune \t count \n")
	for x, v := range counts {
		fmt.Printf("%q \t %d\n", x, v)
	}

	fmt.Printf("nlen \t tcount \n")
	for j, k := range utflen {
		fmt.Printf("%d \t %d \n", j, k)
	}

	if invalid > 0 {
		fmt.Printf("%d invalid UTF-8 characters", invalid)
	}
}
