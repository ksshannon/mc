package eo

import (
	"os"
	"strings"
	"testing"
)

func TestParse1937(t *testing.T) {
	fin, _ := os.Open("data/1937.txt")
	defer fin.Close()
	eos := parseExecOrders(fin)
	if eos == nil {
		t.Fatal("parsing failed")
	}

	// Check the data in the first order
	e := eos[0]
	if e.Number != "7532" {
		t.Errorf("incorrect number: %s", e.Number)
	}
	if strings.Index(e.Title, "Shinnecock") < 0 {
		t.Errorf("incorrect title: %s", e.Title)
	}
	if len(e.Notes) < 1 {
		t.Fatal("invalid notes")
	}
	if n, ok := e.Notes["Revoked by"]; !ok {
		t.Errorf("invalid notes: %+v", e.Notes)
	} else if strings.Index(n, "Public") < 0 {
		t.Errorf("invalid notes: %+v", e.Notes)
	}
}

func TestParse1983(t *testing.T) {
	fin, _ := os.Open("data/1983.txt")
	defer fin.Close()
	eos := parseExecOrders(fin)
	if eos == nil {
		t.Fatal("parsing failed")
	}

	// Find 12407, it should be revoke 12314
	found := false
	for _, e := range eos {
		if e.Number == "12407" {
			found = true
			if strings.Index(e.Notes["Revokes"], "12314") < 0 {
				t.Errorf("invalid revokes note: %s", e.Notes["Revokes"])
			}
		}
	}
	if !found {
		t.Error("couldn't find proper order (12407)")
	}
}

func TestMultiRevoke(t *testing.T) {
	fin, _ := os.Open("data/1979.txt")
	defer fin.Close()
	eos := parseExecOrders(fin)
	if eos == nil {
		t.Fatal("parsing failed")
	}

	found := false
	// Find 12148, revokes many orders, including 10242
	for _, e := range eos {
		if e.Number == "12148" {
			revokes := e.Revokes()
			for _, n := range revokes {
				if n == 10242 {
					found = true
					break
				}
			}
		}
	}
	if !found {
		t.Error("didn't find 10242 in the revoke notes")
	}
}
