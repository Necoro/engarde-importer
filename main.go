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

//goland:noinspection GoUnusedType
type engarde interface {
	engarde() (string, error)
}

type participant struct {
	LastName    string `csv:"lastname"`
	FirstName   string `csv:"firstname"`
	DateOfBirth string `csv:"dateofbirth"`
	Gender      Gender `csv:"gender"`
	Nation      string `csv:"nation"`
	Region      string `csv:"region"`
	Club        string `csv:"club"`
	id          uint
	club        club
}

type club struct {
	name string
	id   uint
}

func (p participant) engarde() (string, error) {
	name := strings.ToUpper(p.LastName)
	gender, err := p.Gender.engarde()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
{[classe tireur] [sexe %s] [presence present] [carton_coach non] [status normal]
 [medical non] [lateralite droite] [nom " %s "] [prenom " %s "]
 [points 1.0] [date_oed "340"] [cle %d] [club1 %d]}
`, gender, name, p.FirstName, p.id, p.club.id), nil
}

func (c club) engarde() (string, error) {
	return fmt.Sprintf(`
{[classe club] [nom "%s"] [modifie vrai] [date_oed "332"] [cle %d]}`, c.name, c.id), nil
}

func prepareParticipants(participants []participant) {
	clubs := map[string]club{}
	var clubId, participantId uint

	for i := range participants {
		p := &participants[i]
		participantId++

		p.id = participantId

		pClub := p.Club
		if strings.Contains(pClub, ", ") {
			pClub = strings.Split(pClub, ", ")[0]
		}

		if c, ok := clubs[pClub]; ok {
			p.club = c
		} else {
			clubId++

			c = club{pClub, clubId}
			p.club = c
			clubs[p.Club] = c
		}
	}
}

func parseOphardtInput(fileName string) ([]participant, error) {
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

	var participants []participant
	if err = dec.Decode(&participants); err != nil {
		return nil, fmt.Errorf("decoding file '%s': %w", fileName, err)
	}

	prepareParticipants(participants)
	return participants, nil
}

func usage() string {
	return fmt.Sprintf("Usage: %s <input csv> <output dir> <name> <M/F> <S/V> <D/F/S> <dd.mm.yyyy>", os.Args[0])
}

type EngardeConfig struct {
	inputFile string
	outputDir string
	name      string
	gender    Gender
	ageGroup  AgeGroup
	weapon    Weapon
	date      time.Time
}

func parseArgs() (config EngardeConfig, err error) {
	config.inputFile = os.Args[1]
	config.outputDir = os.Args[2]
	config.name = os.Args[3]

	if config.gender, err = GenderFromString(os.Args[4]); err != nil {
		return EngardeConfig{}, err
	}
	if config.ageGroup, err = AgeGroupFromString(os.Args[5]); err != nil {
		return EngardeConfig{}, err
	}
	if config.weapon, err = WeaponFromString(os.Args[6]); err != nil {
		return EngardeConfig{}, err
	}

	if config.date, err = time.Parse("02.01.2006", os.Args[7]); err != nil {
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

	_, err = parseOphardtInput(cfg.inputFile)
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
