package comparejson

import (
	"encoding/json"
)

type DBEntries struct {
	List []DBEntry
}
type DBEntry struct {
	Id   string `json: "id"`
	Name string `json: "name"`
}

func EqualDataSets(s1 *DBEntries, s2 *DBEntries) bool {
	if len(s1.List) != len(s2.List) {
		return false
	}

	set1 := ToHashSet(s1.List)
	set2 := ToHashSet(s2.List)

	for k, _ := range set1 {
		if !set2[k] {
			return false
		}
	}

	return true
}

func ToHashSet(list []DBEntry) map[DBEntry]bool {
	hashSet := make(map[DBEntry]bool)
	for _, e := range list {
		hashSet[e] = true
	}
	return hashSet
}

func ParseJsonDataSet(dataJson string) (*DBEntries, error) {
	arr := &DBEntries{}
	err := json.Unmarshal([]byte(dataJson), &arr.List)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
