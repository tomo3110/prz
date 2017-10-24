package container

import (
	"io/ioutil"
	"io"
)

var doc *document

func init() {
	doc = nil
}

type document struct {
	MarkdownBody string
}

func NewDocument() *document {
	if doc == nil {
		doc = &document{""}
	} else {
		doc.MarkdownBody = ""
	}
	return doc
}


func (d *document) SetMarkdownBody(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	d.MarkdownBody = string(b)
	return nil
}
