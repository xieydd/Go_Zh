package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sorter/algorithms/bubblesort"
	"sorter/algorithms/qsort"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithms *string = flag.String("a", "qsort", "Sort algorithms")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file", infile)
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile=", *infile, " outfile=", *outfile, " algorithms=", *algorithms)
	}

	values, err := readValues(*infile)

	if err == nil {
		t1 := time.Now()
		switch *algorithms {
		case "bubblesort":
			bubblesort.BubbleSort(values)
		case "qsort":
			qsort.QSort(values)
		default:
			fmt.Println("Sorting algorithms:", *algorithms, "is either known or expected")
		}
		t2 := time.Now()
		fmt.Println("Sorting process cost ", t2.Sub(t1), "to complete")
		writeValues(values, *outfile)
		fmt.Println("Read values: ", values)
	} else {
		fmt.Println(err)
	}
}
