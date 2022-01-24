package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

func prepareParticipants(participants []Participant) []Club {
	clubMap := map[string]uint{}
	var clubs []Club
	var clubId, participantId uint

	for i := range participants {
		p := &participants[i]
		participantId++

		p.Id = participantId

		pClub := p.ClubStr
		if strings.Contains(pClub, ", ") {
			pClub = strings.Split(pClub, ", ")[0]
		}

		if cId, ok := clubMap[pClub]; ok {
			p.ClubId = cId
		} else {
			clubId++

			p.ClubId = clubId
			clubMap[pClub] = clubId
			clubs = append(clubs, Club{pClub, clubId})
		}
	}

	return clubs
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

func usage() string {
	return fmt.Sprintf("Usage: %s <input csv> <output dir> <name> <M/F> <S/V> <D/F/S> <dd.mm.yyyy>", os.Args[0])
}

type EngardeConfig struct {
	inputFile    string
	outputDir    string
	Name         string
	Description  string
	Gender       Gender
	AgeGroup     AgeGroup
	Weapon       Weapon
	Date         time.Time
	Participants []Participant
	Clubs        []Club
}

func parseArgs() (config EngardeConfig, err error) {
	config.inputFile = os.Args[1]
	config.outputDir = os.Args[2]
	config.Name = os.Args[3]

	if config.Gender, err = GenderFromString(os.Args[4]); err != nil {
		return EngardeConfig{}, err
	}
	if config.AgeGroup, err = AgeGroupFromString(os.Args[5]); err != nil {
		return EngardeConfig{}, err
	}
	if config.Weapon, err = WeaponFromString(os.Args[6]); err != nil {
		return EngardeConfig{}, err
	}

	if config.Date, err = time.Parse("02.01.2006", os.Args[7]); err != nil {
		return EngardeConfig{}, err
	}

	return config, nil
}

func run() error {
	if len(os.Args) <= 7 {
		return errors.New(usage())
	}

	cfg, err := parseArgs()
	if err != nil {
		return err
	}

	cfg.Participants, cfg.Clubs, err = parseOphardtInput(cfg.inputFile)
	if err != nil {
		return err
	}

	return write(cfg)
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("An error occured: %v", err)
	}
}
