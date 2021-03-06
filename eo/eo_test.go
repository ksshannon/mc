// Copyright 2016 Kyle Shannon.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eo

import (
	"testing"
)

func TestWhom(t *testing.T) {
	tests := []struct {
		name string
		i    string
	}{{Unknown, "1"},
		{Roosevelt, "6071"},
		{Roosevelt, "9537"},
		{Truman, "9540"},
		{Carter, "11968"},
		{Reagan, "12300"},
		{BushHW, "12668"},
		{BushHW, "12833"},
		{Clinton, "12944"},
		{Clinton, "12987"},
		{BushW, "13198"},
		{BushW, "13488"},
		{Obama, "13489"},
		{Obama, "13490"},
		{Obama, "13500"},
		{Obama, "13764"},
		{Trump, "13765"},
		{Trump, "20000"},
	}
	var e ExecOrder
	for _, test := range tests {
		eo := &ExecOrder{Number: test.i}
		w := eo.Whom()
		if w != test.name {
			t.Errorf("failed whom: %+v (got %s)", test, w)
		}
		e.Number = test.i
		w = e.Whom()
		if w != test.name {
			t.Errorf("failed Whom: %+v (got %s)", test, w)
		}
	}
}

func TestString(t *testing.T) {
	eo := ExecOrder{
		Number: "9414",
		Title:  "Regulations Relating to Annual and Sick Leave of Government Employees",
		Notes: map[string]string{
			"Signed":                         "January 13, 1944",
			"Federal Register page and date": "9 FR 623, January 18, 1944",
			"Supersedes":                     "EO 8384, March 29, 1940; EO 8385, March 29, 1940; EO 9307, March 3, 1943; EO 9371, August 24, 1943",
			"Note":                           "The authority of this Executive order was repealed by the Annual and Sick Leave Act of 1951.",
		},
	}
	_ = eo
}

func TestNoteKeys(t *testing.T) {
	keys := map[string]int{}
	eos, err := ParseAllOrders("./data")
	if err != nil {
		t.Fatal(err)
	}
	for _, eo := range eos {
		for k := range eo.Notes {
			keys[k]++
		}
	}
	for k, _ := range keys {
		if _, ok := noteTypos[k]; !ok {
			t.Errorf("unknown key: %s (%d)", k, keys[k])
		}
	}
}
