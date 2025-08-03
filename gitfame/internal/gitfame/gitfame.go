package gitfame

import (
	"gitlab.com/slon/shad-go/gitfame/internal/pkg/gitcommand"
)

type AuthorInfo struct {
	Name    string `json:"name"`
	Commits int    `json:"commits"`
	Lines   int    `json:"lines"`
	Files   int    `json:"files"`
}

type AuthorStats struct {
	Files   map[string]struct{}
	Commits map[string]struct{}
	Lines   int
}

func Gitfame(g *gitcommand.GitCommandDescription) ([]AuthorInfo, error) {
	filenames, err := g.GetGitFiles()
	if err != nil {
		return nil, err
	}

	authorStats := make(map[string]*AuthorStats)

	for _, file := range filenames {
		blameRes, err := g.BlameFile(file)
		if err != nil {
			return nil, err
		}

		if len(blameRes) == 0 {
			curCommit, curAuthor, err := g.LogFile(file)
			if err != nil {
				return nil, err
			}
			updateAuthorStats(authorStats, curAuthor.FullName, file, curCommit, 0)
		} else {
			for sha, authorLines := range blameRes {
				updateAuthorStats(authorStats, authorLines.FullName, file, sha, authorLines.Lines)
			}
		}
	}

	authorsList := compileAuthorInfo(authorStats)

	return authorsList, nil
}

func updateAuthorStats(stats map[string]*AuthorStats, author, file, commit string, lines int) {
	if _, exists := stats[author]; !exists {
		stats[author] = &AuthorStats{
			Files:   make(map[string]struct{}),
			Commits: make(map[string]struct{}),
			Lines:   0,
		}
	}
	stats[author].Files[file] = struct{}{}
	stats[author].Commits[commit] = struct{}{}
	stats[author].Lines += lines
}

func compileAuthorInfo(stats map[string]*AuthorStats) []AuthorInfo {
	var authorsList []AuthorInfo
	for author, stat := range stats {
		authorsList = append(authorsList, AuthorInfo{
			Name:    author,
			Commits: len(stat.Commits),
			Lines:   stat.Lines,
			Files:   len(stat.Files),
		})
	}
	return authorsList
}
