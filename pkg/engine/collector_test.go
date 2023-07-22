package engine

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestSourceCollectorAddEntry(t *testing.T) {
	collector := newSourceCollector()
	collector.addEntry("hi", "there")
	assert.True(t, collector.entries["hi"].Contains("there"))
}

func TestSourceCollectorAddPatternMatch(t *testing.T) {
	collector := newSourceCollector()
	pattern := PatternMatch{Kind: FilePatternMatch, Pattern: "test"}
	collector.addPatternMatch([]PatternMatch{pattern})
	assert.True(t, collector.patternMatches.Contains(pattern))
}

func TestSourceCollectorAddPatternExclusion(t *testing.T) {
	collector := newSourceCollector()
	collector.addPatternExclusions([]string{"test"})
	assert.True(t, collector.patternExclusions.Contains("test"))
}

func TestSourceCollectorRun(t *testing.T) {
	tests := []struct {
		patternMatches []PatternMatch
		entries        map[string][]string
		exclusions     map[string][]string
	}{
		{
			[]PatternMatch{
				{
					Kind:    FilePatternMatch,
					Pattern: "file1",
				},
			},
			map[string][]string{"file1": {"testdata/collectorsrc/file1.txt"}},
			nil,
		},
		{
			[]PatternMatch{
				{
					Kind:    ExtensionPatternMatch,
					Pattern: ".txt",
				},
			},
			map[string][]string{
				".txt": {
					"testdata/collectorsrc/file1.txt",
					"testdata/collectorsrc/file2.txt",
					"testdata/collectorsrc/file3.txt",
				},
			},
			nil,
		},
		{
			[]PatternMatch{
				{
					Kind:    DirectoryPatternMatch,
					Pattern: "dir1",
				},
			},
			map[string][]string{
				"dir1": {"testdata/collectorsrc/dir1"},
			},
			nil,
		},
		{
			[]PatternMatch{
				{
					Kind:       ExtensionPatternMatch,
					Pattern:    ".txt",
					Exclusions: &[]string{"file2.txt"},
				},
			},
			map[string][]string{
				".txt": {
					"testdata/collectorsrc/file1.txt",
					"testdata/collectorsrc/file3.txt",
				},
			},
			map[string][]string{
				".txt": {
					"testdata/collectorsrc/file2.txt",
				},
			},
		},
	}
	for _, tc := range tests {
		collector := newSourceCollector()
		collector.patternMatches = mapset.NewSet[PatternMatch](tc.patternMatches...)
		if err := collector.run("testdata/collectorsrc"); err != nil {
			t.Fatal(err)
		}
		for k, v := range tc.entries {
			assert.True(t, collector.entries[k].Contains(v...))
		}
		for k, v := range tc.exclusions {
			assert.False(t, collector.entries[k].Contains(v...))
		}
	}
}
