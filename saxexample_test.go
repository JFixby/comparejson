package comparejson

import (
	"fmt"
	"github.com/jfixby/pin"
	"io"
	"strings"
	"testing"
)

type ElementType string

const ELEMENT = "ELEMENT"
const LIST = "LIST"

func TestExamples(T *testing.T) {

	parseJsonSax(examplejs3)

}

func TestHatchExample(t *testing.T) {

	parseJsonSax(examplejs1)
	parseJsonSax(examplejs2)
}

func parseJsonSax(jsonStream string) {
	var stream io.Reader = strings.NewReader(jsonStream)
	params := &JsonReaderParams{
		stream:   &stream,
		listener: &JsonReaderTestListener{},
	}
	reader := NewJsonReader(params)
	reader.ReadAll()
}

type JsonReaderTestListener struct {
}

func (l *JsonReaderTestListener) OnBeginDocument() {}

func (l *JsonReaderTestListener) OnEndDocument() {}

func (l *JsonReaderTestListener) OnBeginElement(name string, path []string) {
	pin.D(fmt.Sprintf("%v", strings.Join(path[:], "/")), name+" : {")
}
func (l *JsonReaderTestListener) OnEndElement(name string, path []string) {
	pin.D(fmt.Sprintf("%v", strings.Join(path[:], "/")), "} : "+name)
}

func (l *JsonReaderTestListener) OnBeginList(name string, path []string) {
	pin.D(fmt.Sprintf("%v", strings.Join(path[:], "/")), name+" : [")
}

func (l *JsonReaderTestListener) OnEndList(name string, path []string) {
	pin.D(fmt.Sprintf("%v", strings.Join(path[:], "/")), "] : "+name)
}

func (l *JsonReaderTestListener) OnAttribute(key string, value string, path []string) {
	pin.D(fmt.Sprintf("%v", strings.Join(path[:], "/")), value)
}

func (l *JsonReaderTestListener) OnError(e error) {
	pin.E("", e)
}
