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

	parseJson(js1)
	parseJson(js2)
}
