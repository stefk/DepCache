package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Cache struct {
	handlers []Handler
}

type Handler interface {
	Name() string
	Version() (string, error)
	Applicable([]os.FileInfo) (bool, error)
}

func (c Cache) AddHandler(h Handler) {
	version, err := h.Version()

	if err != nil {
		log.Printf("Cannot add handler %q: unable to get underlying version (package manager might not be available)", h.Name())
	} else {
		log.Printf("Adding handler %q (version: %s)", h.Name(), version)
		c.handlers = append(c.handlers, h)
	}
}

func (c Cache) Process(requestDir string) (solvedTar string, err error) {
	log.Printf("Processing directory %q", requestDir)

	files, err := ioutil.ReadDir(requestDir)

	if err != nil {
		return "", err
	}

	hasHandler := false

	for _, h := range c.handlers {
		applicable, err := h.Applicable(files)

		if err != nil {
			return "", err
		}

		if applicable {
			hasHandler = true
			log.Printf("Applying handler %q", h.Name())
		} else {
			log.Printf("Handler %q not applicable to request dir %q", h.Name(), requestDir)
		}
	}

	if !hasHandler {
		return "", fmt.Errorf("No applicable handler for request dir %q", requestDir)
	}

	return "test", nil
}

func main() {
	cache := Cache{}
	cache.AddHandler(Bower{})
	//cache.AddHandler(Composer{})
	//cache.AddHandler(NPM{})
	tar, err := cache.Process("data")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Requested tar = '%s'", tar)
	}
}
