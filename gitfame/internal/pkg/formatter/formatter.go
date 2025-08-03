package formatter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"text/tabwriter"
)

const (
	Tabular   = "tabular"
	CSV       = "csv"
	JSON      = "json"
	JSONLines = "json-lines"
)

type Formatter struct {
	w       io.Writer
	printer func(slice AuthorSort, w io.Writer) error
}

func New(format string, w io.Writer) *Formatter {
	return &Formatter{
		printer: getFunc(format),
		w:       w,
	}
}

func getFunc(format string) func(slice AuthorSort, w io.Writer) error {
	formatters := map[string]func(slice AuthorSort, w io.Writer) error{
		Tabular:   tabular,
		CSV:       csvPrint,
		JSON:      jsonPrint,
		JSONLines: jsonLines,
	}
	if f, ok := formatters[format]; ok {
		return f
	}
	panic("invalid format option")
}

func tabular(a AuthorSort, out io.Writer) error {
	w := tabwriter.NewWriter(out, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Name\tLines\tCommits\tFiles")

	for _, auth := range a.Authors {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\n", auth.Name, auth.Lines, auth.Commits, auth.Files)
	}

	return w.Flush()
}

func csvPrint(a AuthorSort, out io.Writer) error {
	w := csv.NewWriter(out)
	if err := w.Write([]string{"Name", "Lines", "Commits", "Files"}); err != nil {
		return err
	}

	for _, auth := range a.Authors {
		if err := w.Write([]string{
			auth.Name,
			strconv.Itoa(auth.Lines),
			strconv.Itoa(auth.Commits),
			strconv.Itoa(auth.Files),
		}); err != nil {
			return err
		}
	}

	w.Flush()
	return w.Error()
}

func jsonPrint(a AuthorSort, out io.Writer) error {
	b, err := json.Marshal(a.Authors)
	if err != nil {
		return err
	}
	_, err = out.Write(b)
	return err
}

func jsonLines(a AuthorSort, out io.Writer) error {
	for _, auth := range a.Authors {
		b, err := json.Marshal(auth)
		if err != nil {
			return err
		}
		if _, err = out.Write(b); err != nil {
			return err
		}
		if _, err = out.Write([]byte{'\n'}); err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) Print(a AuthorSort) {
	sort.Sort(a)
	if err := f.printer(a, f.w); err != nil {
		fmt.Fprintf(f.w, "error printing authors: %v\n", err)
	}
}
