package comparejson

import (
	"testing"
)

const EQUAL = true

func TestEqualDataSets(t *testing.T) {
	checkDataSets(t, examplejs1, examplejs2, EQUAL)
	checkDataSets(t, examplejs1, examplejs4, !EQUAL)
}

func checkDataSets(t *testing.T, s1 string, s2 string, expected bool) {
	entriesSet1, _ := ParseJsonDataSet(s1)
	entriesSet2, _ := ParseJsonDataSet(s2)

	eq := EqualDataSets(entriesSet1, entriesSet2)

	if eq != expected {
		t.Fatalf("Expected: %v received %v", expected, eq)
		t.FailNow()
	}
}
