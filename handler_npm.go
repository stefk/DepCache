package main

type NPM struct{}

func (NPM) Name() string {
	return "NPM"
}

func (NPM) Version() (string, error) {
	return CmdOutput("npm", "--version")
}
