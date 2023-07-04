package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

const version = "0.0.4"

var revision = "HEAD"

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}

func csvFileName(prefix, sheet, suffix string) string {
	s := prefix + "_" + sheet + "_" + suffix
	return strings.Trim(s, "_") + ".csv"
}

func main() {
	var showVersion bool
	var showHelp bool
	var prefix string
	var suffix string

	// options
	fs := flag.NewFlagSet("excel2csv", flag.ExitOnError)
	fs.BoolVar(&showVersion, "version", false, "show version")
	fs.BoolVar(&showVersion, "v", false, "show version")
	fs.BoolVar(&showHelp, "help", false, "show help")
	fs.BoolVar(&showHelp, "h", false, "show help")
	fs.StringVar(&prefix, "prefix", "", "output file prefix")
	fs.StringVar(&suffix, "suffix", "", "output file suffix")
	fs.Usage = func() {
		fmt.Println(`Usage:
  excel2csv <path>

Flags:`)
		fs.PrintDefaults()
		fmt.Println(`Repository:
  https://github.com/zztkm/excel2csv`)
	}

	fs.Parse(os.Args[1:])

	if showVersion {
		fmt.Printf("version: %s, revision: %s\n", version, revision)
		return
	}
	if showHelp {
		fs.Usage()
		return
	}

	if flag.NArg() != 1 {
		fatal(errors.New("please specify excel file path. `excel2csv -h` for more details"))
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
		file, err := os.Create(csvFileName(prefix, sh, suffix))
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
