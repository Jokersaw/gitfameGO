package gitcommand

import "os/exec"

type GitCommandDescription struct {
	Revision    string
	Directory   string
	UseCommiter bool
	Predicate   func(string) bool
}

type AuthorLines struct {
	FullName string
	Lines    int
}

func (g *GitCommandDescription) execCommand(name string, args ...string) *exec.Cmd {
	command := exec.Command(name, args...)
	command.Dir = g.Directory
	return command
}
