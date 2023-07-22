package engine

import (
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

var collector *SourceCollector = newSourceCollector()

type SourceCollector struct {
	entries           map[string]mapset.Set[string]
	patternMatches    mapset.Set[PatternMatch]
	patternExclusions mapset.Set[string]
}

func (collector *SourceCollector) Search(key string) []string {
	if set, ok := collector.entries[key]; ok {
		return set.ToSlice()
	}
	return nil
}

func (collector *SourceCollector) addEntry(key, value string) {
	if _, ok := collector.entries[key]; !ok {
		collector.entries[key] = mapset.NewSet[string]()
	}
	collector.entries[key].Add(value)
}

func (collector *SourceCollector) addPatternMatch(patterns []PatternMatch) {
	collector.patternMatches.Append(patterns...)
}

func (collector *SourceCollector) addPatternExclusions(patterns []string) {
	collector.patternExclusions.Append(patterns...)
}

func (collector *SourceCollector) run(source string) error {
	if err := filepath.WalkDir(source, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		for _, match := range collector.patternMatches.ToSlice() {
			if match.Exclusions != nil {
				for _, exclusion := range *match.Exclusions {
					if strings.Contains(path, exclusion) {
						return nil
					}
				}
			}
			name := info.Name()
			if match.Kind == DirectoryPatternMatch && !info.IsDir() {
				continue
			}
			re, err := regexp.Compile(match.Pattern)
			if err != nil {
				return err
			}
			if re.MatchString(name) {
				collector.addEntry(match.Pattern, path)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func newSourceCollector() *SourceCollector {
	return &SourceCollector{
		entries:           make(map[string]mapset.Set[string]),
		patternMatches:    mapset.NewSet[PatternMatch](),
		patternExclusions: mapset.NewSet[string](),
	}
}

type PatternMatchKind int

const (
	FilePatternMatch PatternMatchKind = iota
	DirectoryPatternMatch
	ExtensionPatternMatch
)

type PatternMatch struct {
	Kind       PatternMatchKind
	Key        string
	Pattern    string
	Matched    bool
	Exclusions *[]string
}

func RegisterPatternMatches(patterns []PatternMatch) {
	collector.addPatternMatch(patterns)
}

func RegisterPatternExclusions(patterns []string) {
	collector.addPatternExclusions(patterns)
}

var globalPatternExlusions = []string{
	"node_modules",
}

func init() {
	RegisterPatternExclusions(globalPatternExlusions)
}
