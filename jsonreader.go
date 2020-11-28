package comparejson

import (
	"encoding/json"
	"fmt"
	"io"
)

type JsonReader interface {
	ReadAll()
}

type JsonReaderListener interface {
	OnBeginDocument()
	OnEndDocument()

	OnBeginElement(name string, path []string)
	OnEndElement(name string, path []string)

	OnBeginList(name string, path []string)
	OnEndList(name string, path []string)

	OnAttribute(key string, value string, path []string)

	OnError(e error)
}

type simpleJsonReader struct {
	stream   *io.Reader
	decoder  *json.Decoder
	listener JsonReaderListener
}

type readJsonArgs struct {
	decoder  *json.Decoder
	listener JsonReaderListener
	path     []string
}

func (args *readJsonArgs) lastKey() string {
	return args.path[len(args.path)-1]
}

func (args *readJsonArgs) push(token string) {
	args.path = append(args.path, token)
}

func (args *readJsonArgs) pop() {
	args.path = args.path[0 : len(args.path)-1]
}

func (s *simpleJsonReader) ReadAll() {
	args := &readJsonArgs{
		decoder:  s.decoder,
		listener: s.listener,
		path:     []string{},
	}
	readJson(args)
}

func readJson(args *readJsonArgs) error {
	args.listener.OnBeginDocument()
	args.push("")
	_, err := readJsonValue(args)
	if err != nil {
		args.listener.OnError(err)
		return err
	}
	args.pop()
	args.listener.OnEndDocument()
	return nil
}

const ExitScope = true

func readJsonValue(args *readJsonArgs) (bool, error) {
	token, err := readToken(args)

	if err != nil {
		return ExitScope, err
	}

	if token == "{" {
		args.listener.OnBeginElement(args.lastKey(), args.path)
		readElement(args)

		return !ExitScope, nil
	}

	if token == "]" {
		args.listener.OnEndList(args.lastKey(), args.path)
		return ExitScope, nil
	}
	if token == "[" {
		args.listener.OnBeginList(args.lastKey(), args.path)
		readList(args)

		return !ExitScope, nil
	}

	args.listener.OnAttribute(args.lastKey(), token, args.path)

	return !ExitScope, nil
}

func readList(args *readJsonArgs) error {
	i := 0
	for {
		args.push("")
		i++
		exitScope, err := readJsonValue(args)
		if err != nil {
			return err
		}
		args.pop()
		if exitScope {
			return nil
		}
	}
}

func readElement(args *readJsonArgs) error {
	for {
		token, err := readToken(args)
		if err != nil {
			return err
		}

		if token == "}" {
			args.listener.OnEndElement(args.lastKey(), args.path)
			return nil
		}
		args.push(token)
		exitScope, err := readJsonValue(args)
		if err != nil {
			return err
		}
		args.pop()
		if exitScope {
			return nil
		}
	}
}

func readToken(args *readJsonArgs) (string, error) {
	entry, err := args.decoder.Token()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", entry), nil
}

type JsonReaderParams struct {
	stream   *io.Reader
	listener JsonReaderListener
}

func NewJsonReader(params *JsonReaderParams) JsonReader {
	decoder := json.NewDecoder(*params.stream)
	return &simpleJsonReader{
		stream:   params.stream,
		decoder:  decoder,
		listener: params.listener,
	}
}
