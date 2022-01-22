package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jszwec/csvutil"
)

type Participant struct {
	LastName    string `csv:"lastname"`
	FirstName   string `csv:"firstname"`
	DateOfBirth string `csv:"dateofbirth"`
	Gender      string `csv:"gender"`
	Nation      string `csv:"nation"`
	Region      string `csv:"region"`
	Club        string `csv:"club"`
}

func parseOphardtInput(fileName string) ([]Participant, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("opening input file '%s': %w", fileName, err)
	}

	encReader, err := getEncodedReader(f)
	if err != nil {
		return nil, fmt.Errorf("cannot determine encoding of file '%s': %w", fileName, err)
	}

	csvReader := csv.NewReader(encReader)
	csvReader.Comma = ';'

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, fmt.Errorf("reading from file '%s': %w", fileName, err)
	}
	dec.DisallowMissingColumns = true

	var participants []Participant
	if err = dec.Decode(&participants); err != nil {
		return nil, fmt.Errorf("decoding file '%s': %w", fileName, err)
	}

	return participants, nil
}

func run() error {
	if len(os.Args) <= 1 {
		return errors.New("need filename to start with")
	}
	fileName := os.Args[1]

	p, err := parseOphardtInput(fileName)
	if err != nil {
		return err
	}

	fmt.Print(p)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("An error occured: %v", err)
	}
}
