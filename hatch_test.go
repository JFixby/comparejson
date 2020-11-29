package comparejson

import (
	"encoding/json"
	"testing"
)

const EQUAL = true

func TestUnmarshall(t *testing.T) {
	checkDataSets(t, examplejs1, examplejs2, EQUAL)
	checkDataSets(t, examplejs1, examplejs4, !EQUAL)
}

func checkDataSets(t *testing.T, s1 string, s2 string, expected bool) {
	entriesSet1, _ := parseJsonSet(s1)
	entriesSet2, _ := parseJsonSet(s2)

	eq := EqualDataSets(entriesSet1, entriesSet2)

	if eq != expected {
		t.Fatalf("Expected: %v received %v", expected, eq)
		t.FailNow()
	}
}

func EqualDataSets(s1 *DBEntries, s2 *DBEntries) bool {
	if len(s1.List) != len(s2.List) {
		return false
	}

	set1 := toSet(s1.List)
	set2 := toSet(s2.List)

	for k, _ := range set1 {
		if !set2[k] {
			return false
		}
	}

	return true
}

func toSet(list []DBEntry) map[DBEntry]bool {
	hashSet := make(map[DBEntry]bool)
	for _, e := range list {
		hashSet[e] = true
	}
	return hashSet
}

type DBEntry struct {
	Id   string `json: "id"`
	Name string `json: "name"`
}

type DBEntries struct {
	List []DBEntry
}

func parseJsonSet(dataJson string) (*DBEntries, error) {
	arr := &DBEntries{}
	err := json.Unmarshal([]byte(dataJson), &arr.List)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
