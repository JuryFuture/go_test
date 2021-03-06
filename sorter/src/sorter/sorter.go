package main

import (
	"flag"
	"bufio"
	"io"
	"os"
	"strconv"
	"fmt"
	"time"
	"algorithms/qsort"
	"algorithms/bubblesort"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Faild to open the input file", infile)
		return
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
			fmt.Println("A too long line, seems unexpected.")
		}

		str := string(line)

		value , err2 := strconv.Atoi(str)

		if err2 != nil {
			err = err2
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	for _,value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)

	if err == nil {
		t1 := time.Now()
		fmt.Println(err)
		switch *algorithm {
			case "qsort":
				qsort.QuickSort(values)
			case "bubblesort":
				bubblesort.BubbleSort(values)
			default:
				fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsuportted")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

	    writeValues(values, *outfile)
	} else {
		fmt.Println("Read values:",values)
	}
}
