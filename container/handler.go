package container

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type handler struct {
	sync.Once
	filename string
	doc      *document
	tmpl     *template.Template
}

func NewHandler(filename string, doc *document) *handler {
	return &handler{filename: filename, doc: doc}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Once.Do(func() {
		h.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", h.filename)))
	})
	if err := h.tmpl.Execute(w, h.doc); err != nil {
		log.Fatal(err)
	}
}
