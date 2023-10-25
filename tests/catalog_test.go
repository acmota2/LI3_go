package tests

import (
	"bufio"
	"os"
	"testing"

	"LI3_go/catalog"
	parser "LI3_go/csv_parser"
	st "LI3_go/single_types"
)

var expected = st.Commit{
	Repo_id:      32294845,
	Author_id:    11392088,
	Committer_id: 11392088,
	Commit_at:    "2015-03-15 19:10:25",
	Message:      "Merge pull request #2 from SimhaGitHub/readme-edits		Update README.md",
}

func setup[K comparable, T any](ft func([]string) T, kt func(*T) K) catalog.Catalog[K, T] {
	f, _ := os.Open("./commits_test.csv")
	defer f.Close()
	bufrd := bufio.NewReader(f)

	return catalog.New(parser.ConvertData(bufrd), ft, kt)
}

func TestMakeCatalog(t *testing.T) {
	c := setup(st.CreateCommit, st.CommitKeyFunc)

	if !expected.Equal(c.View[1]) {
		t.Errorf("Expected %s, got %s", expected, *c.View[1])
	}
}

func TestCatalogHash(t *testing.T) {
	c := setup(st.CreateCommit, st.CommitKeyFunc)

	current := c.Hash[st.CommitKeyFunc(&expected)]
	if !expected.Equal(current) {
		t.Errorf("Expected %s, got %s", expected, *current)
	}
}
