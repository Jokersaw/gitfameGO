package formatter

import (
	"strings"

	"gitlab.com/slon/shad-go/gitfame/internal/gitfame"
)

type AuthorSort struct {
	Authors []gitfame.AuthorInfo
	OrderBy string
}

func (a AuthorSort) Len() int {
	return len(a.Authors)
}

func (a AuthorSort) Swap(i, j int) {
	a.Authors[i], a.Authors[j] = a.Authors[j], a.Authors[i]
}

func less(f, s []int, def bool) bool {
	for i := 0; i < len(f) && i < len(s); i++ {
		if f[i] != s[i] {
			return f[i] > s[i]
		}
	}
	return def
}

func (a AuthorSort) Less(i, j int) bool {
	var a1, a2 []int

	switch a.OrderBy {
	case "lines":
		a1 = []int{a.Authors[i].Lines, a.Authors[i].Commits, a.Authors[i].Files}
		a2 = []int{a.Authors[j].Lines, a.Authors[j].Commits, a.Authors[j].Files}
	case "files":
		a1 = []int{a.Authors[i].Files, a.Authors[i].Lines, a.Authors[i].Commits}
		a2 = []int{a.Authors[j].Files, a.Authors[j].Lines, a.Authors[j].Commits}
	case "commits":
		a1 = []int{a.Authors[i].Commits, a.Authors[i].Lines, a.Authors[i].Files}
		a2 = []int{a.Authors[j].Commits, a.Authors[j].Lines, a.Authors[j].Files}
	default:
		panic("bad sort order")
	}

	return less(a1, a2, strings.ToLower(a.Authors[i].Name) < strings.ToLower(a.Authors[j].Name))
}
