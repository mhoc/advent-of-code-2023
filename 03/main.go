package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type EngineSchematic struct {
	Raw      string
	RawLines []string
}

func ReadEngineSchematic() EngineSchematic {
	f, _ := os.ReadFile("./data.txt")
	return EngineSchematic{
		Raw:      string(f),
		RawLines: strings.Split(string(f), "\n"),
	}
}

func (es EngineSchematic) CandidatePartNumbers() []CandidatePartNumber {
	ps := []CandidatePartNumber{}
	for lineNumber, line := range es.RawLines {
		numStr := ""
		for cPosition, c := range line {
			if unicode.IsDigit(c) {
				numStr += string(c)
			}
			if len(numStr) > 0 && (!unicode.IsDigit(c) || cPosition == len(line)-1) {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				var position int
				if unicode.IsDigit(c) {
					position = cPosition - len(numStr) + 1
				} else {
					position = cPosition - len(numStr)
				}
				ps = append(ps, CandidatePartNumber{
					LineNumber: lineNumber,
					Position:   position,
					Value:      num,
					ValueStr:   numStr,
				})
				numStr = ""
			}
		}
	}
	return ps
}

func (es EngineSchematic) Gears() []Gear {
	gs := []Gear{}
	for lineNumber, line := range es.RawLines {
		for cPosition, c := range line {
			if c == '*' {
				gs = append(gs, Gear{
					LineNumber: lineNumber,
					Position:   cPosition,
				})
			}
		}
	}
	return gs
}

func (es EngineSchematic) Line(i int) string {
	return es.RawLines[i]
}

type CandidatePartNumber struct {
	LineNumber int
	Position   int
	Value      int
	ValueStr   string
}

func (cpn CandidatePartNumber) HasSymbolBottom(es EngineSchematic) bool {
	if cpn.LineNumber == len(es.RawLines)-1 {
		return false
	}
	nextLine := es.Line(cpn.LineNumber + 1)
	start := cpn.Position
	if start != 0 {
		// Covers the bottom-left diagonal
		start -= 1
	}
	end := cpn.Position + len(cpn.ValueStr)
	if end < len(nextLine) {
		// Covers the bottom-right diagonal
		end += 1
	}
	for _, c := range nextLine[start:end] {
		if !unicode.IsDigit(c) && c != '.' {
			return true
		}
	}
	return false
}

func (cpn CandidatePartNumber) HasSymbolLeft(es EngineSchematic) bool {
	if cpn.Position == 0 {
		return false
	}
	leftChar := es.Line(cpn.LineNumber)[cpn.Position-1]
	if !unicode.IsDigit(rune(leftChar)) && leftChar != '.' {
		return true
	}
	return false
}

func (cpn CandidatePartNumber) HasSymbolRight(es EngineSchematic) bool {
	line := es.Line(cpn.LineNumber)
	if cpn.Position+len(cpn.ValueStr) > len(line)-1 {
		return false
	}
	rightChar := line[cpn.Position+len(cpn.ValueStr)]
	if !unicode.IsDigit(rune(rightChar)) && rightChar != '.' {
		return true
	}
	return false
}

func (cpn CandidatePartNumber) HasSymbolTop(es EngineSchematic) bool {
	if cpn.LineNumber == 0 {
		return false
	}
	precedingLine := es.Line(cpn.LineNumber - 1)
	start := cpn.Position
	if start != 0 {
		// covers the top left case
		start -= 1
	}
	end := cpn.Position + len(cpn.ValueStr)
	if end < len(precedingLine) {
		// covers the top right case
		end += 1
	}
	for _, c := range precedingLine[start:end] {
		if !unicode.IsDigit(c) && c != '.' {
			return true
		}
	}
	return false
}

func (cpn CandidatePartNumber) IsPartNumber(es EngineSchematic) bool {
	return cpn.HasSymbolTop(es) ||
		cpn.HasSymbolRight(es) ||
		cpn.HasSymbolBottom(es) ||
		cpn.HasSymbolLeft(es)
}

type Gear struct {
	LineNumber int
	Position   int
}

func (g Gear) PerimeterNumbers(es EngineSchematic) []int {
	ns := []int{}
	// TOP
	if g.LineNumber != 0 {
		prevLine := es.Line(g.LineNumber - 1)
		if unicode.IsDigit(rune(prevLine[g.Position])) {
			// Check Top Center first. If there's a digit here, its the only thing we need to check for top, because it would
			// inherently cover the top-left and top-right cases.
			t := FindNumberWithin(prevLine, g.Position)
			if t != -1 {
				ns = append(ns, t)
			}
		} else {
			tl := FindNumberWithin(prevLine, g.Position-1)
			if tl != -1 {
				ns = append(ns, tl)
			}
			tr := FindNumberWithin(prevLine, g.Position+1)
			if tr != -1 {
				ns = append(ns, tr)
			}
		}
	}
	// Right
	{
		r := FindNumberWithin(es.Line(g.LineNumber), g.Position+1)
		if r != -1 {
			ns = append(ns, r)
		}
	}
	// BOTTOM
	if g.LineNumber != len(es.RawLines)-1 {
		nextLine := es.Line(g.LineNumber + 1)
		if unicode.IsDigit(rune(nextLine[g.Position])) {
			b := FindNumberWithin(nextLine, g.Position)
			if b != -1 {
				ns = append(ns, b)
			}
		} else {
			bl := FindNumberWithin(nextLine, g.Position-1)
			if bl != -1 {
				ns = append(ns, bl)
			}
			br := FindNumberWithin(nextLine, g.Position+1)
			if br != -1 {
				ns = append(ns, br)
			}
		}
	}
	// Left
	{
		l := FindNumberWithin(es.Line(g.LineNumber), g.Position-1)
		if l != -1 {
			ns = append(ns, l)
		}
	}
	return ns
}

func FindNumberWithin(line string, anchoredAt int) int {
	if len(line) == 0 || anchoredAt < 0 || anchoredAt > len(line)-1 {
		return -1
	}
	if !unicode.IsDigit(rune(line[anchoredAt])) {
		return -1
	}
	start, end := anchoredAt, anchoredAt+1
	for {
		if start-1 >= 0 && unicode.IsDigit(rune(line[start-1])) {
			start -= 1
		} else {
			break
		}
	}
	for {
		if end < len(line) && unicode.IsDigit(rune(line[end])) {
			end += 1
		} else {
			break
		}
	}
	v, err := strconv.Atoi(line[start:end])
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	es := ReadEngineSchematic()
	cpns := es.CandidatePartNumbers()
	var total int
	for _, cpn := range cpns {
		if cpn.IsPartNumber(es) {
			total += cpn.Value
		}
	}
	log.Printf("Parts: %v", total)
	gears := es.Gears()
	var totalGearRatio int
	for _, gear := range gears {
		fmt.Printf("%+v ", gear)
		gearPerimeter := gear.PerimeterNumbers(es)
		fmt.Printf("%v\n", gearPerimeter)
		if len(gearPerimeter) == 2 {
			gearRatio := 1
			for _, perimeterNumber := range gearPerimeter {
				gearRatio *= perimeterNumber
			}
			totalGearRatio += gearRatio
		}
	}
	log.Printf("Gear Ratios: %v", totalGearRatio)
}
