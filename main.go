package main

import (
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"fmt"
)

// CurrentURL is the header of the column for the current landing page
const CurrentURL = "Current Landing Page"

// RedirectURL is the header for the column that pages will be redirected to
const RedirectURL = "Redirect Landing Page"

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Error program must be envoked with ", os.Args[0], " <xlsx> <output>")
	}
	redirects := ReadXlsx(os.Args[1])
	WriteFile(os.Args[2], redirects)
}

// WriteFile writes the redirects to a file. Note, this will need to be copy pasted from the file to your config file.
func WriteFile(fileName string, redirects []Redirect) {
	fh, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("could not create file: ", err)
	}
	defer fh.Close()
	for _, redirect := range redirects {
		fh.WriteString(redirect.String())
	}
	fh.Sync()
}

// ReadXlsx reads the XLSX for the redirect URLs
func ReadXlsx(fileName string) []Redirect {
	xls, err := xlsx.OpenFile(fileName)

	if err != nil {
		log.Fatal("Can't open file error: ", err)
	}

	sheet := xls.Sheets[0]
	redirects := make([]Redirect, len(sheet.Rows))
	redirect, current := getHeaderIndex(sheet.Rows[0])
	for index, row := range sheet.Rows[1:] {
		// XLSXs like to leave trailing rows.
		if row.Cells[redirect].Value == "" && row.Cells[current].Value == "" {
			return redirects[:index]
		}
		redirects[index].RedirectURL = row.Cells[redirect].Value
		redirects[index].CurrentURL = row.Cells[current].Value
	}

	return redirects
}

// getHeaderIndex is a helper to find the proper headers in the XLSX
func getHeaderIndex(header *xlsx.Row) (int, int) {
	redirect, current := -1, -1
	for index, cell := range header.Cells {
		if cell.Value == CurrentURL {
			current = index
		}
		if cell.Value == RedirectURL {
			redirect = index
		}
	}
	if redirect == -1 || current == -1 {
		panic("Inproper headers in XLSX file, requires " + CurrentURL + " and " + RedirectURL)
	}

	return redirect, current
}

// A structure which contains the component parts of a redirect
type Redirect struct {
	CurrentURL string
	RedirectURL string
}

func (r Redirect) String() string {
	return fmt.Sprintf("rewrite ^%s$ \t %s permanent;\n", r.CurrentURL, r.RedirectURL)
}