package sqltgo

import (
	"bytes"
	"text/template"
)

type SQLT struct {
	Template *template.Template
	SqlRaw   string
}

func NewSQLT() *SQLT {
	return &SQLT{}
}

func NewParse(name, text string, data interface{}) (*SQLT, error) {
	t, err := NewSQLT().Parse(name, text)
	if err != nil {
		return nil, err
	}
	results, err := t.Execute(data)
	if err != nil {
		return nil, err
	}
	t.SqlRaw = results
	return t, nil
}

func (s *SQLT) ParseFiles(filenames ...string) (*SQLT, error) {
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		return nil, err
	}
	s.Template = t
	return s, nil
}

func (s *SQLT) Parse(name, text string) (*SQLT, error) {
	t, err := template.New(name).Parse(text)
	if err != nil {
		return nil, err
	}
	s.Template = t
	return s, nil
}

func (s *SQLT) Execute(data interface{}) (string, error) {
	var b bytes.Buffer
	err := s.Template.Execute(&b, data)
	if err != nil {
		return "", err
	}
	s.SqlRaw = b.String()
	return s.SqlRaw, nil
}
