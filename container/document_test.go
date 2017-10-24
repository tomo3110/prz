package container

import (
	"bytes"
	"testing"
	"fmt"
)

func TestSetMarkdownBody(t *testing.T) {
	const testText = "# Test"
	f := bytes.NewBufferString(testText)
	doc := NewDocument(f)
	if err := doc.SetMarkdownBody(f); err != nil {
		t.Error(err)
	}
	if doc.MarkdownBody != testText {
		t.Fatalf("expected: %s to: %s", testText, doc.MarkdownBody)
	}
}
