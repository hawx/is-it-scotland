package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hawx/is-it-scotland/trie"
)

func isItScotland(code string) string {
	switch code {
	case "E92000001":
		return "no"
	case "S92000003":
		return "yes"
	case "W92000004":
		return "no"
	}

	return "-"
}

func readInputCsv() (*trie.Trie, error) {
	t := trie.New()

	r := csv.NewReader(os.Stdin)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		t.Add(strings.ReplaceAll(row[0], " ", ""), isItScotland(row[4]))
	}

	return t, nil
}

func readDataset(path string) (*trie.Trie, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	t := trie.New()

	r := csv.NewReader(file)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		t.Add(row[0], row[1])
	}

	return t, nil
}

func doBuild() {
	t, err := readInputCsv()
	if err != nil {
		log.Println(err)
	}

	t.Optimise()

	w := csv.NewWriter(os.Stdout)
	for key, value := range t.AsMap() {
		w.Write([]string{key, value})
	}
	w.Flush()
}

func main() {
	build := flag.Bool("build", false, "print optimised dataset from input CSVs")
	dataset := flag.String("dataset", "", "dataset to query")
	flag.Parse()

	if *build {
		doBuild()
	} else {
		t, err := readDataset(*dataset)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Enter postcode to query, CTRL+C to quit:")

		r := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("\n > ")
			line, _ := r.ReadString('\n')
			query := strings.ReplaceAll(strings.ToUpper(string(line)), " ", "")

			fmt.Printf("=> %s\n", t.Get(query))
		}
	}
}
