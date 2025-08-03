package filter

import (
	"path"
	"strings"
)

type Filter struct {
	Extensions []string
	Languages  []string
	Exclude    []string
	RestrictTo []string

	extensions map[string]struct{}
	langExt    map[string]struct{}
}

func (f *Filter) initializeMaps() {
	if len(f.extensions) == 0 && len(f.Extensions) > 0 {
		f.extensions = make(map[string]struct{})
		for _, ext := range f.Extensions {
			f.extensions[ext] = struct{}{}
		}
	}

	if len(f.langExt) == 0 && len(f.Languages) > 0 {
		f.langExt = make(map[string]struct{})
		for _, l := range f.Languages {
			for _, ext := range nameToLanguage[strings.ToLower(l)].Extensions {
				f.langExt[ext] = struct{}{}
			}
		}
	}
}

func (f *Filter) isExcluded(filename string) bool {
	for _, exclude := range f.Exclude {
		if match, err := path.Match(exclude, filename); err == nil && match {
			return true
		}
	}
	return false
}

func (f *Filter) isRestricted(filename string) bool {
	if len(f.RestrictTo) == 0 {
		return true
	}

	for _, glob := range f.RestrictTo {
		if match, err := path.Match(glob, filename); err == nil && match {
			return true
		}
	}
	return false
}

func (f *Filter) hasValidExtension(filename string) bool {
	ext := path.Ext(filename)
	if len(f.extensions) > 0 {
		if _, ok := f.extensions[ext]; !ok {
			return false
		}
	}
	if len(f.langExt) > 0 {
		if _, ok := f.langExt[ext]; !ok {
			return false
		}
	}
	return true
}

func (f *Filter) DoMatch(filename string) bool {
	f.initializeMaps()

	if f.isExcluded(filename) {
		return false
	}

	if !f.isRestricted(filename) {
		return false
	}

	if !f.hasValidExtension(filename) {
		return false
	}

	return true
}
