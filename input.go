package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
)

type Participant struct {
	LastName    string `csv:"lastname"`
	FirstName   string `csv:"firstname"`
	DateOfBirth string `csv:"dateofbirth"`
	Gender      Gender `csv:"gender"`
	Nation      string `csv:"nation"`
	Region      string `csv:"region"`
	ClubStr     string `csv:"club"`
	Id          uint   `csv:"-"`
	ClubId      uint   `csv:"-"`
}

type Club struct {
	Name string
	Id   uint
}

func (p Participant) MainClub() string {
	if strings.Contains(p.ClubStr, ", ") {
		return strings.Split(p.ClubStr, ", ")[0]
	}
	return p.ClubStr
}

func parseOphardtInput(fileName string) ([]Participant, []Club, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("opening input file '%s': %w", fileName, err)
	}

	encReader, err := encodedReader(f)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot determine encoding of file '%s': %w", fileName, err)
	}

	csvReader := csv.NewReader(encReader)
	csvReader.Comma = ';'

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, nil, fmt.Errorf("reading from file '%s': %w", fileName, err)
	}
	dec.DisallowMissingColumns = true

	var participants []Participant
	if err = dec.Decode(&participants); err != nil {
		return nil, nil, fmt.Errorf("decoding file '%s': %w", fileName, err)
	}

	clubs := prepareParticipants(participants)
	return participants, clubs, nil
}

func prepareParticipants(participants []Participant) []Club {
	clubMap := map[string]uint{}
	var clubs []Club
	var clubId, participantId uint

	for i := range participants {
		participantId++

		p := &participants[i]
		p.Id = participantId

		club := p.MainClub()
		if cId, ok := clubMap[club]; ok {
			p.ClubId = cId
		} else {
			clubId++

			p.ClubId = clubId
			clubMap[club] = clubId
			clubs = append(clubs, Club{club, clubId})
		}
	}

	return clubs
}
