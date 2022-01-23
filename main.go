package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
)

type enguarde interface {
	enguarde() (string, error)
}

type Gender int

const (
	GenderM = iota
	GenderF
)

func (g Gender) String() string {
	switch g {
	case GenderM:
		return "M"
	case GenderF:
		return "F"
	default:
		return fmt.Sprintf("U%d", g)
	}
}

func (g Gender) enguarde() (string, error) {
	switch g {
	case GenderM:
		return "masculin", nil
	case GenderF:
		return "feminin", nil
	default:
		return "", fmt.Errorf("unknown gender value '%d'", g)
	}
}

func (g *Gender) UnmarshalCSV(content []byte) error {
	c := string(content)
	switch c {
	case "M":
		*g = GenderM
	case "F":
		*g = GenderF
	default:
		return fmt.Errorf("unknown gender value '%s'", c)
	}

	return nil
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

func (p participant) enguarde() (string, error) {
	name := strings.ToUpper(p.LastName)
	gender, err := p.Gender.enguarde()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
{[classe tireur] [sexe %s] [presence present] [carton_coach non] [status normal]
 [medical non] [lateralite droite] [nom " %s "] [prenom " %s "]
 [points 1.0] [date_oed "340"] [cle %d] [club1 %d]}
`, gender, name, p.FirstName, p.id, p.club.id), nil
}

func (c club) enguarde() (string, error) {
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

func run() error {
	if len(os.Args) <= 1 {
		return errors.New("need filename to start with")
	}
	fileName := os.Args[1]

	p, err := parseOphardtInput(fileName)
	if err != nil {
		return err
	}

	x, _ := p[0].enguarde()
	y, _ := p[1].enguarde()

	fmt.Print(x, y)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("An error occured: %v", err)
	}
}
