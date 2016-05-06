package main

import (
	"os"
)

type Bower struct{}

func (Bower) Name() string {
	return "Bower"
}

func (Bower) Version() (string, error) {
	return CmdOutput("bower", "--version")
}

func (Bower) Applicable([]os.FileInfo) (bool, error) {
	return true, nil
}
