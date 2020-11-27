package comparejson

import (
	"testing"
)

type TestExpectedScenario string

const ERROR = "ERROR"
const TRUE = "TRUE"
const FALSE = "FALSE"

func TestHatch(t *testing.T) {

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
	js2 := `[
	{
	"id":"wqweq",
	"name":"test json 2" },
	{
	"name":"test json",
	"id":"jhasdad"
	}
]`

	ChechEquals(t, js1, js2, TRUE)

}

func ChechEquals(t *testing.T, js1 string, js2 string, expected TestExpectedScenario) {
	result, err := AreEqualJSON(js1, js2)

	outcome := interpretTestResult(result, err)
	if err != nil {
		t.Fatalf(err.Error())
		t.FailNow()
	}
	if outcome != expected {
		t.Fatalf("Test failed: expected (%v), received (%v)", expected, outcome)
		t.FailNow()
	}

}

func interpretTestResult(result bool, err error) TestExpectedScenario {
	if err != nil {
		return ERROR
	}
	if result {
		return TRUE
	}
	return FALSE
}
