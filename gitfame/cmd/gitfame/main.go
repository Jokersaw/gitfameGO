//go:build !solution

package main

import (
	"os"

	"gitlab.com/slon/shad-go/gitfame/internal/flags"
	"gitlab.com/slon/shad-go/gitfame/internal/gitfame"
	"gitlab.com/slon/shad-go/gitfame/internal/pkg/filter"
	"gitlab.com/slon/shad-go/gitfame/internal/pkg/formatter"
	"gitlab.com/slon/shad-go/gitfame/internal/pkg/gitcommand"
)

func main() {
	flagsInfo := flags.GetFlags(os.Args[1:])

	filter := &filter.Filter{
		Extensions: flagsInfo.ExtentionsFlag,
		Languages:  flagsInfo.LanguagesFlag,
		Exclude:    flagsInfo.ExcludeFlag,
		RestrictTo: flagsInfo.RestrictToFlag,
	}

	gitcommand := &gitcommand.GitCommandDescription{
		Revision:    flagsInfo.RevisionFlag,
		Directory:   flagsInfo.RepositoryFlag,
		UseCommiter: flagsInfo.UseCommitterFlag,
		Predicate:   filter.DoMatch,
	}

	table, err := gitfame.Gitfame(gitcommand)

	if err != nil {
		panic(err)
	}

	authorSlice := &formatter.AuthorSort{
		Authors: table,
		OrderBy: flagsInfo.OrderByFlag,
	}

	formatter := formatter.New(flagsInfo.FormatFlag, os.Stdout)

	formatter.Print(*authorSlice)
}
