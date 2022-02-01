package main

import "fmt"

type Gender int32

const (
	GenderM Gender = iota
	GenderF
)

var GenderStrings = []string{"Herren", "Damen"}

func (g Gender) String() string {
	return GenderStrings[g]
}

func (g Gender) Engarde() (string, error) {
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

type AgeGroup int32

const (
	AgeVeteran AgeGroup = iota
	AgeSenior
)

var AgeGroupStrings = []string{"Veteranen", "Senioren"}

func (a AgeGroup) String() string {
	return AgeGroupStrings[a]
}

func (a AgeGroup) Engarde() (string, error) {
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

type Weapon int32

const (
	Epee Weapon = iota
	Foil
	Sabre
)

var WeaponStrings = []string{"Degen", "Florett", "Säbel"}
var WeaponShorts = []string{"D", "F", "S"}

func (w Weapon) String() string {
	return WeaponStrings[w]
}

func (w Weapon) ShortString() string {
	return WeaponShorts[w]
}

func (w Weapon) Engarde() (string, error) {
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
