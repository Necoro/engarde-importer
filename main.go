package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type EngardeConfig struct {
	InputFile    string
	OutputDir    string
	Name         string
	Description  string
	Gender       Gender
	AgeGroup     AgeGroup
	Weapon       Weapon
	Date         time.Time
	Participants []Participant
	Clubs        []Club
}

func usage() string {
	return fmt.Sprintf("Usage: %s <input csv> <output dir> <name> <M/F> <S/V> <D/F/S> <dd.mm.yyyy>", os.Args[0])
}

func parseArgs() (config EngardeConfig, err error) {
	config.InputFile = os.Args[1]
	config.OutputDir = os.Args[2]
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

	cfg.Participants, cfg.Clubs, err = parseOphardtInput(cfg.InputFile)
	if err != nil {
		return err
	}

	return Write(cfg)
}

func main() {
	if len(os.Args) == 1 {
		gui()
	} else {
		if err := run(); err != nil {
			log.Fatalf("An error occured: %v", err)
		}
	}
}
