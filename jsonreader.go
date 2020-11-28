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
	OnBegin()
	OnEnd()

	OnBeginElement(name string, level int)
	OnEndElement(name string)

	OnBeginList(name string, level int)
	OnAddListElement()
	OnEndList(name string)

	OnAttribute(key string, value string)

	OnError(e error)
}

type simpleJsonReader struct {
	stream   *io.Reader
	decoder  *json.Decoder
	listener JsonReaderListener
}

func (s *simpleJsonReader) ReadAll() {
	readJson(s.decoder, s.listener)
}

func readJson(d *json.Decoder, listener JsonReaderListener) error {
	listener.OnBegin()
	_, err := readJsonValue(d, "", listener, 0)
	if err != nil {
		listener.OnError(err)
		return err
	}
	listener.OnEnd()
	return nil
}

const ExitScope = true

func readJsonValue(d *json.Decoder, key string, listener JsonReaderListener, level int) (bool, error) {
	token, err := readToken(d)

	if err != nil {
		return ExitScope, err
	}

	if token == "{" {
		//pin.D("begin element", key)
		listener.OnBeginElement(key, level+1)
		readElement(d, key, listener, level+1)

		return !ExitScope, nil
	}

	if token == "]" {
		//pin.D("end array", key)
		listener.OnEndList(key)
		return ExitScope, nil
	}
	if token == "[" {
		//pin.D("begin array", key)
		listener.OnBeginList(key, level+1)
		readList(d, key, listener, level+1)

		return !ExitScope, nil
	}

	//pin.D(key, token)
	listener.OnAttribute(key, token)

	return !ExitScope, nil
}

func readList(d *json.Decoder, key string, listener JsonReaderListener, level int) error {
	for {
		exit_scope, err := readJsonValue(d, "", listener, level)
		if err != nil {
			return err
		}
		if exit_scope {
			return nil
		}

		listener.OnAddListElement()
	}
}

func readElement(d *json.Decoder, key string, listener JsonReaderListener, level int) error {
	for {
		token, err := readToken(d)
		if err != nil {
			return err
		}
		if token == "}" {
			//pin.D("end element", key)
			listener.OnEndElement(key)
			return nil
		}
		exit_scope, err := readJsonValue(d, token, listener, level)
		if err != nil {
			return err
		}
		if exit_scope {
			return nil
		}
	}
}

func readToken(d *json.Decoder) (string, error) {
	entry, err := d.Token()
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
