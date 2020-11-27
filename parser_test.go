package comparejson

import (
	"fmt"
	"github.com/jfixby/pin"
	"testing"
)

type ElementType string

const ELEMENT = "ELEMENT"
const LIST = "LIST"

func TestParser(t *testing.T) {

	js1 := `[
	{
		"id":"jhasdad",
		"name":"test json"
	},
	{
		"id":"wqweq",
		"name":"test json 2"
	}
]`

	error := parse(js1)
	if error != nil {
		t.Fatalf(error.Error())
		t.FailNow()
	}

}

func parse(s string) error {
	error := parseSlice(s, 0, len(s))
	return error
}

func parseSlice(s string, from int, to int) error {
	oindex, oetype := indexOfOpening(s, from, to)
	cindex, cetype := indexOfClosing(s, from, to)
	if oetype != cetype {
		return fmt.Errorf(
			"invalid json element opens with (%v):%v and closes with (%v):%v",
			oetype, oindex, cetype, cindex)
	}
	if oetype == "" {
		return nil
	}

	pin.D("body", s[oindex:cindex])

	return parseSlice(s, oindex+1, cindex-1)
}

func indexOfOpening(s string, from int, to int) (int, ElementType) {
	for i := from; i < to; i++ {
		e := s[i : i+1]
		if e == "{" {
			return i, ELEMENT
		}
		if e == "[" {
			return i, LIST
		}
	}
	return -1, ""
}

func indexOfClosing(s string, from int, to int) (int, ElementType) {
	for i := to; i >= from; i-- {
		e := s[i-1 : i]
		if e == "}" {
			return i, ELEMENT
		}
		if e == "]" {
			return i, LIST
		}
	}
	return -1, ""
}
