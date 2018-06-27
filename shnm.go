//Sensible Human Name Masking Tool
// What the tool should do: Produce logical names and ulid's using name and genders

package main

import (
	"bytes"
	cryptorand "crypto/rand"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	mathrand "math/rand"
	"os"
	"strconv"
	"time"

	"./pkg/nameutils"
	"github.com/oklog/ulid"
)

type vertex struct {
	f, l int
}

var generated map[vertex]bool

func generate() (int, int) {
	//If you seed a source with the same number, it produces the same sequence of random numbers.
	s1 := mathrand.NewSource(time.Now().UnixNano())
	r1 := mathrand.New(s1)
	for {
		i := r1.Intn(4999) + 1
		j := r1.Intn(4999) + 1
		v := vertex{i, j}
		if !generated[v] {
			generated[v] = true
			return i, j
		}
	}
}

func printusage() {
	fmt.Fprintf(os.Stderr, "SHNM produces sensible human names along with masking information.\r\n"+
		"For example, one could effectively mask patient, next-of-kin, clinican names with sensible sounding names\r\n"+
		"and provide corelated informaion with the mask so that data could be re-identified in the future.")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = printusage
	numsamples := flag.Int("n", 1, "Number of human names to be generated (max 25 million)")
	gender := flag.String("g", "", "Gender of the human names to be generated 'f' for females and 'm' for males. will generate female if unspecified")
	genUlids := flag.Bool("u", false, "Generate ULID's (false by default), set to true if you want them")
	flag.Parse()
	if *numsamples > 25000000 {
		fmt.Fprintf(os.Stderr, "Too many samples requested or unspecified input\r\n")
		os.Exit(1)
	}

	if !(*gender == "m" || *gender == "f" || *gender == "") {
		fmt.Fprintf(os.Stderr, "Cannot understand the gender requested - unspecified input\r\n")
		os.Exit(1)
	}

	// If ULIDs are required then check if we can generate them
	entropy := cryptorand.Reader
	if *genUlids == true {
		_, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}

	//Create a map to store the vertexes we are generating
	generated = make(map[vertex]bool)

	// Get Names list
	//Get data from the internal store - you can generate this using go-bindata
	csvFile, err := nameutils.Asset("data/20180626_human_names.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	r := csv.NewReader(bytes.NewReader(csvFile))
	records, _ := r.ReadAll()

	// Set the gender preference
	cnum := 1
	if *gender == "m" {
		cnum = 2
	}

	w := csv.NewWriter(os.Stdout)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	for len(generated) < *numsamples {
		firstname, surname := generate()
		if *genUlids == true {
			id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			w.Write([]string{strconv.Itoa(firstname), strconv.Itoa(surname), records[firstname][cnum], records[surname][3], id.String()})
		} else {
			w.Write([]string{strconv.Itoa(firstname), strconv.Itoa(surname), records[firstname][cnum], records[surname][3]})

		}
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
