package main

import "fmt"

type Gender int

const (
	GenderM Gender = iota
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

func GenderFromString(content string) (Gender, error) {
	switch content {
	case "M":
		return GenderM, nil
	case "F":
		return GenderF, nil
	default:
		return 0, fmt.Errorf("unknown gender value '%s'", content)
	}
}

func (g *Gender) UnmarshalCSV(content []byte) error {
	if res, err := GenderFromString(string(content)); err == nil {
		*g = res
		return nil
	} else {
		return err
	}
}

type AgeGroup int

const (
	AgeVeteran AgeGroup = iota
	AgeSenior
)

func (a AgeGroup) String() string {
	switch a {
	case AgeVeteran:
		return "V"
	case AgeSenior:
		return "S"
	default:
		return fmt.Sprintf("U%d", a)
	}
}

func (a AgeGroup) enguarde() (string, error) {
	switch a {
	case AgeVeteran:
		return "veteran", nil
	case AgeSenior:
		return "senior", nil
	default:
		return "", fmt.Errorf("unknown age group value '%d'", a)
	}
}

func AgeGroupFromString(content string) (AgeGroup, error) {
	switch content {
	case "V":
		return AgeVeteran, nil
	case "S":
		return AgeSenior, nil
	default:
		return 0, fmt.Errorf("unknown age group value '%s'", content)
	}
}

type Weapon int

const (
	Epee Weapon = iota
	Foil
	Sabre
)

func (w Weapon) String() string {
	switch w {
	case Epee:
		return "D"
	case Foil:
		return "F"
	case Sabre:
		return "S"
	default:
		return fmt.Sprintf("U%d", w)
	}
}

func (w Weapon) enguarde() (string, error) {
	switch w {
	case Epee:
		return "epee", nil
	case Foil:
		return "fleuret", nil
	case Sabre:
		return "sabre", nil
	default:
		return "", fmt.Errorf("unknown weapon value '%d'", w)
	}
}

func WeaponFromString(content string) (Weapon, error) {
	switch content {
	case "D":
		return Epee, nil
	case "S":
		return Sabre, nil
	case "F":
		return Foil, nil
	default:
		return 0, fmt.Errorf("unknown weapon value '%s'", content)
	}
}
