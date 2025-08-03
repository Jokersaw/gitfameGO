package gitcommand

func (g *GitCommandDescription) LogFile(filename string) (string, *AuthorLines, error) {
	res, err := g.execCommand("git", "log", g.Revision, "--format=%H%an", "-n", "1", "--", filename).Output()

	if err != nil {
		return "", nil, err
	}

	sha := string(res[:40])
	author := res[40 : len(res)-1]

	return sha, &AuthorLines{
		FullName: string(author),
		Lines:    0,
	}, nil
}
