package main

import (
	"strings"
	"testing"
)

var RawEngineSchematic1 = []string{
	`.*123...`,
	`.456..-.`,
	`......80`,
}

func TestReadEngineSchematic(t *testing.T) {
	ReadEngineSchematic()
}

func TestCandidatePartNumbers(t *testing.T) {
	es := EngineSchematic{
		Raw:      strings.Join(RawEngineSchematic1, "\n"),
		RawLines: RawEngineSchematic1,
	}
	pn := es.CandidatePartNumbers()
	if len(pn) != 3 {
		t.Errorf("expected length of 3; got %v", len(pn))
	}
	if pn[0].Value != 123 {
		t.Errorf("expected 0 value of 123; got %v", pn[0].Value)
	}
	if pn[0].ValueStr != "123" {
		t.Errorf("expected 0 valueStr of 123; got %v", pn[0].ValueStr)
	}
	if pn[0].LineNumber != 0 {
		t.Errorf("expected 0 line number of 0; got %v", pn[0].LineNumber)
	}
	if pn[0].Position != 2 {
		t.Errorf("expected 0 position of 2; got %v", pn[0].Position)
	}
	if pn[1].Value != 456 {
		t.Errorf("expected 1 value of 456; got %v", pn[1].Value)
	}
	if pn[1].ValueStr != "456" {
		t.Errorf("expected 1 valueStr of 123; got %v", pn[1].ValueStr)
	}
	if pn[1].LineNumber != 1 {
		t.Errorf("expected 1 line number of 1; got %v", pn[1].LineNumber)
	}
	if pn[1].Position != 1 {
		t.Errorf("expected 1 position of 1; got %v", pn[1].Position)
	}
	if pn[2].Value != 80 {
		t.Errorf("expected 2 value of 80; got %v", pn[2].Value)
	}
	if pn[2].ValueStr != "80" {
		t.Errorf("expected 2 valueStr of 80; got %v", pn[2].ValueStr)
	}
	if pn[2].LineNumber != 2 {
		t.Errorf("expected 2 line number of 2; got %v", pn[2].LineNumber)
	}
	if pn[2].Position != 6 {
		t.Errorf("expected 2 position of 6; got %v", pn[2].Position)
	}
}

func TestCandidatePartNumberHasSymbolBottom(t *testing.T) {
	es := EngineSchematic{
		Raw:      strings.Join(RawEngineSchematic1, "\n"),
		RawLines: RawEngineSchematic1,
	}
	cpns := es.CandidatePartNumbers()
	if cpns[0].HasSymbolBottom(es) {
		t.Errorf("expected 123 to not have a symbol below it, but it does")
	}
	if cpns[1].HasSymbolBottom(es) {
		t.Errorf("expected 456 to not have a symbol below it, but it does")
	}
	if cpns[2].HasSymbolBottom(es) {
		t.Errorf("expected 8 to not have a symbol below it, but it does")
	}
}

func TestCandidatePartNumberHasSymbolLeft(t *testing.T) {
	es := EngineSchematic{
		Raw:      strings.Join(RawEngineSchematic1, "\n"),
		RawLines: RawEngineSchematic1,
	}
	cpns := es.CandidatePartNumbers()
	if !cpns[0].HasSymbolLeft(es) {
		t.Errorf("expected 123 to have a symbol left, but it doesn't")
	}
	if cpns[1].HasSymbolLeft(es) {
		t.Errorf("expected 456 to not have a symbol left, but it does")
	}
	if cpns[2].HasSymbolLeft(es) {
		t.Errorf("expected 8 to not have a symbol left, but it does")
	}
}

func TestCandidatePartNumberHasSymbolRight(t *testing.T) {
	es := EngineSchematic{
		Raw:      strings.Join(RawEngineSchematic1, "\n"),
		RawLines: RawEngineSchematic1,
	}
	cpns := es.CandidatePartNumbers()
	if cpns[0].HasSymbolRight(es) {
		t.Errorf("expected 123 to not have a symbol right, but it does")
	}
	if cpns[1].HasSymbolRight(es) {
		t.Errorf("expected 456 to not have a symbol right, but it does")
	}
	if cpns[2].HasSymbolRight(es) {
		t.Errorf("expected 8 to not have a symbol right, but it does")
	}
}

func TestCandidatePartNumberHasSymbolTop(t *testing.T) {
	es := EngineSchematic{
		Raw:      strings.Join(RawEngineSchematic1, "\n"),
		RawLines: RawEngineSchematic1,
	}
	cpns := es.CandidatePartNumbers()
	if cpns[0].HasSymbolTop(es) {
		t.Errorf("expected 123 to have no symbol above it, but it does")
	}
	if !cpns[1].HasSymbolTop(es) {
		t.Errorf("expected 456 to have a symbol above it, but it doesn't")
	}
	if !cpns[2].HasSymbolTop(es) {
		t.Errorf("expected 8 to have a symbol above it, but it doesn't")
	}
}

func TestFindNumberWithin1(t *testing.T) {
	if FindNumberWithin("", 0) != -1 {
		t.Errorf("does not handle empty strings")
	}
}

func TestFindNumberWithin2(t *testing.T) {
	if FindNumberWithin("..2", 5) != -1 {
		t.Errorf("does not handle out of bound indexes")
	}
}

func TestFindNumberWithin3(t *testing.T) {
	if FindNumberWithin(".22", 0) != -1 {
		t.Errorf("expected -1; got %v", FindNumberWithin(".22", 0))
	}
}

func TestFindNumberWithin4(t *testing.T) {
	if FindNumberWithin(".22", 1) != 22 {
		t.Errorf("expected 22; got %v", FindNumberWithin(".22", 1))
	}
}

func TestFindNumberWithin5(t *testing.T) {
	s := ".4...235"
	v1 := FindNumberWithin(s, 1)
	if v1 != 4 {
		t.Errorf("1")
	}
	v2 := FindNumberWithin(s, 5)
	if v2 != 235 {
		t.Errorf("2")
	}
	v3 := FindNumberWithin(s, 6)
	if v3 != 235 {
		t.Errorf("3")
	}
	v4 := FindNumberWithin(s, 6)
	if v4 != 235 {
		t.Errorf("4")
	}
}
