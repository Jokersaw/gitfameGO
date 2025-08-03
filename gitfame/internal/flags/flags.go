package flags

import (
	"github.com/spf13/pflag"
)

type FlagsInfo struct {
	RepositoryFlag   string
	RevisionFlag     string
	OrderByFlag      string
	UseCommitterFlag bool
	FormatFlag       string
	ExtentionsFlag   []string
	LanguagesFlag    []string
	ExcludeFlag      []string
	RestrictToFlag   []string
}

func GetFlags(args []string) *FlagsInfo {

	set := pflag.NewFlagSet("gitfame", pflag.PanicOnError)
	set.String("repository", ".", "path to the Git repository")
	set.String("revision", "HEAD", "commit revision")
	set.String("order-by", "lines", "order by field: lines, commits, files")
	set.Bool("use-committer", false, "use committer instead of author")
	set.String("format", "tabular", "output format: tabular, csv, json, json-lines")
	set.StringSlice("extensions", make([]string, 0), "file extensions to include")
	set.StringSlice("languages", make([]string, 0), "languages to include")
	set.StringSlice("exclude", make([]string, 0), "glob patterns to exclude files")
	set.StringSlice("restrict-to", make([]string, 0), "glob patterns to include files")
	err := set.Parse(args)

	if err != nil {
		panic(err)
	}

	repository, _ := set.GetString("repository")
	revision, _ := set.GetString("revision")
	orderBy, _ := set.GetString("order-by")
	useCommitter, _ := set.GetBool("use-committer")
	format, _ := set.GetString("format")
	extensions, _ := set.GetStringSlice("extensions")
	languages, _ := set.GetStringSlice("languages")
	exclude, _ := set.GetStringSlice("exclude")
	restrictTo, _ := set.GetStringSlice("restrict-to")

	return &FlagsInfo{
		repository, revision, orderBy, useCommitter, format, extensions, languages, exclude, restrictTo,
	}
}
