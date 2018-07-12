//Clinical Text Spell Checker
//What the tool does: Uses a large clinical dictionary to check for spelling mistakes in a large csv file
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/VertisPro/fasthealth-tools/pkg/textutils"
)

/*
"Database support is provided through the GO SQL (https://github.com/go-sql-driver/mysql) and mattn(https://github.com/mattn/go-sqlite) sqlite library and provides supports for MySQL (4.1+), MariaDB, Percona Server, Google CloudSQL or Sphinx (2.2.3+) and sqlite.\r\n"+
"https://github.com/lib/pq.\r\n"+
*/
func printusage() {
	fmt.Fprintf(os.Stderr, "This tool carries out spelling checks on blobs of clinical text.\r\n"+
		"Blobs are typically supplied through a csv file or individually\r\n"+
		"The tools requires a dictionary file.\r\n"+
		"A large dictionary of (US/AU) medical and (US) english words has been provided.\r\n"+
		"Spelling suggestions are provided as an extra column in the csv file for easy correlation.\r\n"+
		"This tool uses the ENCHANT library (2.2.1) and is provided as part of the FastHealth open source tools under the MIT license.\r\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	//Set up the command line usage
	flag.Usage = printusage
	flag.Parse()
	spellchecker, err := textutils.NewSpellChecker("./data/combined.dic")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	sentence := "The patient came with complaints of chwst paun."
	suggestions, err := spellchecker.CheckSentence(sentence)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(strings.Join(suggestions, " "))
}
