package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type Handler interface {
	Name() string
	Version() (string, error)
}

type Status struct {
	Name    string
	Version string
}

type Bower struct{}

func (Bower) Name() string {
	return "Bower"
}

func (Bower) Version() (string, error) {
	return Output("bower", "--version")
}

type Npm struct{}

func (Npm) Name() string {
	return "NPM"
}

func (Npm) Version() (string, error) {
	return Output("npm", "--version")
}

type Composer struct{}

func (Composer) Name() string {
	return "Composer"
}

func (Composer) Version() (string, error) {
	return Output("composer", "--version")
}

func Output(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).Output()
	return strings.TrimSpace(string(out)), err
}

func HandlerStatus(h Handler) Status {
	v, err := h.Version()
	if err != nil {
		v = "Not available (" + err.Error() + ")"
	}
	return Status{
		Name:    h.Name(),
		Version: v,
	}
}

func main() {
	handlers := [24]Handler{
		Npm{},
		Bower{},
		Composer{},
		Npm{},
		Bower{},
		Composer{},
		Npm{},
		Bower{},
		Composer{},
		Npm{},
		Bower{},
		Composer{},
		Npm{},
		Bower{},
		Composer{},
		Npm{},
		Bower{},
		Composer{},
		Npm{},
		Bower{},
		Composer{},
		Npm{},
		Bower{},
		Composer{},
	}
	for _, h := range handlers {
		s := HandlerStatus(h)
		fmt.Println("- Handler: ", s.Name)
		fmt.Println("  Version: ", s.Version)
	}
}
