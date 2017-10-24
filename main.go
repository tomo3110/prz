package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/skratchdot/open-golang/open"
	"github.com/tomo3110/prz/container"
)

var portStr string

const helpMsg = `Usage of %s:
	%s [OPTIONS] MarkdownFilePaht
Options
`

func run() error {
	if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			return err
		}
		doc := container.NewDocument()
		if err := doc.SetMarkdownBody(f); err != nil {
			return err
		}

		revealHandler := container.NewHandler("index.html", doc)

		open.Run("http://localhost" + portStr + "/slide")

		http.Handle("/", http.FileServer(Assets))
		http.Handle("/slide", revealHandler)
		http.ListenAndServe(portStr, nil)

		return nil
	}
	return errors.New("none Argment: MarkdownFilePath")
}

func main() {
	// flag init
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpMsg, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&portStr, "p", ":8080", "server port conf")
	flag.Parse()

	// running app
	if err := run(); err != nil {
		fmt.Println(err)
		flag.Usage()
		log.Fatal()
	}
	os.Exit(0)
}
