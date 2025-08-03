package gitcommand

import (
	"strings"
)

func (g *GitCommandDescription) GetGitFiles() ([]string, error) {
	result, err := g.execCommand("git", "ls-tree", "-r", "--name-only", g.Revision).Output()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return []string{}, nil
	}

	filenames := strings.Split(string(result[:len(result)-1]), "\n")
	filteredNames := make([]string, 0, len(filenames))

	for _, file := range filenames {
		if g.Predicate(file) {
			filteredNames = append(filteredNames, file)
		}
	}
	return filteredNames, nil
}
