package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VertisPro/fasthealth-tools/pkg/fileutils"
)

func printusage() {
	fmt.Fprintf(os.Stderr, "This tool encrypts and decrypts a file using the supplied passphrase.\r\n"+
		"PBKDF2 is used for generating keys and the file is encrytped using AES-128.\r\n"+
		"There is currently no consensus for an encryted file format so the utility implements its own.\r\n"+
		"The tool is meant for automated/batch usage and does not prompt when over-writing files. \r\n"+
		"It is provided as part of the FastHealth open source tools under the MIT license.\r\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	//Set up the command line usage
	flag.Usage = printusage
	passphrase := flag.String("p", "", "(Required) Password for encrypting or decrypting the file")
	infile := flag.String("i", "", "(Required) Input file")
	outfile := flag.String("o", "", "(Required) Output file")
	decrypt := flag.Bool("d", false, "Decrypt - by defualt the utility will only encrypt the file")
	flag.Parse()

	// Check if all required command line flags are set
	if *passphrase == "" || *infile == "" || *outfile == "" {
		fmt.Fprintf(os.Stderr, "Missing required parameters, please check your input.\r\n")
		printusage()
	}

	// Check if input file exists
	if _, err := os.Stat(*infile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "The input file does not exist or the path is incorrect\r\n")
		os.Exit(1)
	}

	// Check if output file can be written
	f, err := os.Create(*outfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to write an output file, please check path or permissions\r\n")
		os.Exit(1)
	}
	f.Close()

	if *decrypt == false {
		fmt.Println("Encrypting", *infile, "to", *outfile)
		fileutils.AESEncryptFile(*infile, *passphrase, *outfile)
	} else {
		fmt.Println("Decrypting", *infile, "to", *outfile)
		fileutils.AESDecryptFile(*infile, *passphrase, *outfile)
	}
}
