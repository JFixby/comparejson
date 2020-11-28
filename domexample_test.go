package comparejson

import (
	"github.com/jfixby/pin"
	"io"
	"strings"
	"testing"
)

func TestDomExamples(T *testing.T) {

	parseJsonDom(examplejs3)

}

type JsonReaderDom struct {
	Root        *Node
	currentNode *Node
	currentList List
}

func parseJsonDom(jsonStream string) {
	domReader := &JsonReaderDom{}
	var stream io.Reader = strings.NewReader(jsonStream)
	params := &JsonReaderParams{
		stream:   &stream,
		listener: domReader,
	}
	reader := NewJsonReader(params)
	reader.ReadAll()
}

func (l *JsonReaderDom) OnBeginDocument() {
	l.Root = &Node{}
	l.currentNode = l.Root
}

func (l *JsonReaderDom) OnEndDocument() {}

func (l *JsonReaderDom) OnBeginElement(name string, path []string) {
	if l.currentList == nil {
		child := &Node{Name: name}
		child.Parent = l.currentNode
		l.currentNode.Children.Add(name,child)
		l.currentNode = child
	}
}
func (l *JsonReaderDom) OnEndElement(name string, path []string) {
	l.currentNode = l.currentNode.Parent
}

func (l *JsonReaderDom) OnBeginList(name string, path []string) {
	//pin.D(fmt.Sprintf("%v", strings.Join(path[:], "/")), name+" : [")
	list := NewNodeSet()
	l.currentNode.Lists[name] = list
	l.currentList = list

}

func (l *JsonReaderDom) OnEndList(name string, path []string) {
	//pin.D(fmt.Sprintf("%v", strings.Join(path[:], "/")), "] : "+name)
	l.currentList = nil
}

func (l *JsonReaderDom) OnAttribute(key string, value string, path []string) {
	//pin.D(fmt.Sprintf("%v", strings.Join(path[:], "/")), value)
	l.currentNode.Attributes[key] = value
}

func (l *JsonReaderDom) OnError(e error) {
	pin.E("", e)
	panic(e)
}
