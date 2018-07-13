//Clinical Text Spell Checker
//What the tool does: Uses a large clinical dictionary to check for spelling mistakes in a large csv file
package main

import (
	"bufio"
	"flag"
	"io"
	"log"

	"encoding/csv"
	"fmt"
	"os"

	"github.com/VertisPro/fasthealth-tools/pkg/textutils"
)

/*
"Database support is provided through the GO SQL (https://github.com/go-sql-driver/mysql) and mattn(https://github.com/mattn/go-sqlite) sqlite library and provides supports for MySQL (4.1+), MariaDB, Percona Server, Google CloudSQL or Sphinx (2.2.3+) and sqlite.\r\n"+
"https://github.com/lib/pq.\r\n"+
*/
func printusage() {
	fmt.Fprintf(os.Stderr, "\r\nClinical Text Spell Checker\r\nThis tool carries out spelling checks on blobs of clinical text. The input is a csv file along with information on which column contains the text blob. An output file is created containing the original blob and Spelling suggestions for easy correlation. . The tools requires a dictionary file and a large dictionary of (US/AU) medical and (US) english words has been provided. This tool uses the ENCHANT library (2.2.1) and is provided as part of the FastHealth open source tools.\r\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	//Set up the command line usage
	flag.Usage = printusage
	infile := flag.String("i", "", "(Required) Input CSV file")
	outfile := flag.String("o", "", "(Required) Output CSV file, overwritten if exists")
	dicfile := flag.String("d", "ctscwords.dic", "(Optional) External dictionary file")
	textcolumn := flag.Int("c", 1, "(optional) Column where the text is placed (defualt is 1 which means first column) ")
	flag.Parse()

	// Check if input file exists
	if _, err := os.Stat(*dicfile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "The dictionary file does not exist or the path is incorrect\r\n")
		os.Exit(1)
	}

	// Check if input file exists
	if _, err := os.Stat(*infile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "The input file does not exist or the path is incorrect\r\n")
		os.Exit(1)
	}

	if *outfile == "" {
		fmt.Fprintf(os.Stderr, "The outfile has not been specified\r\n")
		os.Exit(1)
	}

	spellchecker, err := textutils.NewSpellChecker(*dicfile)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	//Open input file
	csvFile, _ := os.Open(*infile)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	//setup writer for outout
	csvOut, err := os.Create(*outfile)

	if err != nil {
		log.Fatal("Unable to open output")
		os.Exit(1)
	}
	defer csvOut.Close()
	w := csv.NewWriter(csvOut)

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
			os.Exit(1)
		}
		sentence := line[*textcolumn]
		suggestions, err := spellchecker.CheckSentence(sentence)
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}
		//strings.Join(suggestions, " ")
		if err = w.Write(suggestions); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
