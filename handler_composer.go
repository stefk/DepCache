package main

type Composer struct{}

func (Composer) Name() string {
	return "Composer"
}

func (Composer) Version() (string, error) {
	return CmdOutput("composer", "--version")
}
