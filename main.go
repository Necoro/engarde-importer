package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

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

func usage() string {
	return fmt.Sprintf("Usage: %s <input csv> <output dir> <name> <M/F> <S/V> <D/F/S> <dd.mm.yyyy>", os.Args[0])
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
