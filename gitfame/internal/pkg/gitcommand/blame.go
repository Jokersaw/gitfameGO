package gitcommand

import (
	"strings"
)

func (g *GitCommandDescription) BlameFile(filename string) (map[string]*AuthorLines, error) {
	res, err := g.execCommand("git", "blame", filename, "--porcelain", "-c", g.Revision).Output()
	if err != nil {
		return nil, err
	}

	commitToAuthorLines := make(map[string]*AuthorLines)

	lines := strings.Split(string(res), "\n")
	for i := 0; i < len(lines)-1; {
		curLine := lines[i]
		words := strings.Fields(curLine)

		sha := words[0]
		if _, exists := commitToAuthorLines[sha]; !exists {
			if !g.UseCommiter {
				commitToAuthorLines[sha] = &AuthorLines{
					FullName: lines[i+1][7:],
				}
			} else {
				commitToAuthorLines[sha] = &AuthorLines{
					FullName: lines[i+5][10:],
				}
			}
			if lines[i+10][:8] != "filename" {
				i++
			}
			i += 12
		} else {
			i += 2
		}

		commitToAuthorLines[sha].Lines++
	}

	return commitToAuthorLines, nil
}
