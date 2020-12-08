package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	birthYear      = "byr"
	issueYear      = "iyr"
	expirationYear = "eyr"
	height         = "hgt"
	hairColor      = "hcl"
	eyeColor       = "ecl"
	passportID     = "pid"
	countryID      = "cid"
)

var (
	recordSeparator = []byte("\n\n")
)

// Passport represents a passport
type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      int
}

// CheckValidity checks if the passport is valid
func (p *Passport) CheckValidity() error {
	if p.BirthYear < 1920 || p.BirthYear > 2002 {
		return fmt.Errorf("invalid BirthYear: %d", p.BirthYear)
	}

	if p.IssueYear < 2010 || p.IssueYear > 2020 {
		return fmt.Errorf("invalid IssueYear: %d", p.IssueYear)
	}

	if p.ExpirationYear < 2020 || p.ExpirationYear > 2030 {
		return fmt.Errorf("invalid ExpirationYear: %d", p.ExpirationYear)
	}

	if strings.HasSuffix(p.Height, "cm") {
		offset := strings.Index(p.Height, "cm")
		h, err := strconv.Atoi(p.Height[:offset])
		if err != nil {
			return fmt.Errorf("invalid Height: %s", err)
		}
		if h < 150 || h > 193 {
			return fmt.Errorf("invalid Height: %dcm", h)
		}
	} else if strings.HasSuffix(p.Height, "in") {
		offset := strings.Index(p.Height, "in")
		h, err := strconv.Atoi(p.Height[:offset])
		if err != nil {
			return fmt.Errorf("invalid Height: %s", err)
		}
		if h < 59 || h > 76 {
			return fmt.Errorf("invalid Height: %din", h)
		}
	} else {
		return fmt.Errorf("invali Height: %s", p.Height)
	}

	if hairColorOK, err := regexp.MatchString(`#[0-9a-f]{6}`, p.HairColor); err != nil || !hairColorOK {
		return fmt.Errorf("invalid HairColor: %s", p.HairColor)
	}

	switch p.EyeColor {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
		break
	default:
		return fmt.Errorf("invalid EyeColor: %s", p.EyeColor)
	}

	if len(p.PassportID) != 9 {
		return fmt.Errorf("invalid PassportID: %s", p.PassportID)
	}
	pid, err := strconv.Atoi(p.PassportID)
	if err != nil || pid <= 0 {
		return fmt.Errorf("invalid PassportID: %s", p.PassportID)
	}

	// if p.CountryID <= 0 {
	return nil
}

// IsValid returns true if passport is valid
func (p *Passport) IsValid() bool {
	return p.CheckValidity() == nil
}

// PassportFromString loads a passport from a byte slice
func PassportFromString(data string) (Passport, error) {
	passport := Passport{}
	var err error

	for _, field := range strings.Fields(data) {
		pair := strings.Split(field, ":")
		switch pair[0] {
		case birthYear:
			passport.BirthYear, err = strconv.Atoi(pair[1])
			if err != nil {
				return passport, err
			}
			break
		case issueYear:
			passport.IssueYear, err = strconv.Atoi(pair[1])
			if err != nil {
				return passport, err
			}
			break
		case expirationYear:
			passport.ExpirationYear, err = strconv.Atoi(pair[1])
			if err != nil {
				return passport, err
			}
			break
		case height:
			passport.Height = pair[1]
			break
		case hairColor:
			passport.HairColor = pair[1]
			break
		case eyeColor:
			passport.EyeColor = pair[1]
			break
		case passportID:
			passport.PassportID = pair[1]
			break
		case countryID:
			passport.CountryID, err = strconv.Atoi(pair[1])
			if err != nil {
				return passport, err
			}
			break
		default:
			return passport, fmt.Errorf("unexpected field: %s", pair[0])
		}
	}

	err = passport.CheckValidity()
	if err != nil {
		return passport, err
	}

	return passport, nil
}

func passportSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, recordSeparator); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

// PassportBlocksFromReader loads passport blocks from an io.Reader
func PassportBlocksFromReader(r io.Reader) []string {
	blocks := []string{}
	scanner := bufio.NewScanner(r)

	scanner.Split(passportSplitter)

	for scanner.Scan() {
		block := scanner.Text()

		if strings.Contains(block, birthYear+":") &&
			strings.Contains(block, issueYear+":") &&
			strings.Contains(block, expirationYear+":") &&
			strings.Contains(block, height+":") &&
			strings.Contains(block, hairColor+":") &&
			strings.Contains(block, eyeColor+":") &&
			strings.Contains(block, passportID+":") {

			blocks = append(blocks, block)
		}
	}

	return blocks
}

// PassportBlocksFromFile loads passports from a file
func PassportBlocksFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return PassportBlocksFromReader(file), nil
}

// Part1 solves part 1 of todays challenge
func Part1(blocks []string) int {
	return len(blocks)
}

// Part2 solved part 2 of todays challenge
func Part2(blocks []string) int {
	valid := 0

	for _, block := range blocks {
		passport, err := PassportFromString(block)
		if err != nil || !passport.IsValid() {
			// fmt.Fprintln(os.Stderr, "Problem with passport:", err)
			continue
		}

		valid++
	}

	return valid
}

func main() {
	fmt.Println("Day 4")
	fmt.Println("=====")

	blocks, err := PassportBlocksFromFile("testdata/day4")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file", err)
		os.Exit(1)
	}

	fmt.Println("Part 1:", Part1(blocks))
	fmt.Println("Part 2:", Part2(blocks))
}
