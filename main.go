package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

const version = "0.0.1"

var revision = "HEAD"

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func main() {
	var showVersion bool

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Printf("version: %s, revision: %s\n", version, revision)
		return
	}

	if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}

	filename := flag.Arg(0)

	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fatal(err)
		}
	}()
	// Get all the rows in the Sheet1.
	for _, sh := range f.GetSheetList() {
		file, err := os.Create(sh + ".csv")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		cw := csv.NewWriter(file)
		defer cw.Flush()

		rows, err := f.Rows(sh)
		if err != nil {
			fatal(err)
		}
		for rows.Next() {
			row, err := rows.Columns()
			if err != nil {
				fmt.Println(err)
			}
			cw.Write(row)
		}
		if err = rows.Close(); err != nil {
			fmt.Println(err)
		}

	}
}
