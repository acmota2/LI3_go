package tests

import (
	"bufio"
	"os"
	"testing"

	"LI3_go/catalog"
	parser "LI3_go/csv_parser"
	st "LI3_go/single_types"
)

var ExpectedCommit = st.Commit{
	Repo_id:      32294845,
	Author_id:    11392088,
	Committer_id: 11392088,
	Commit_at:    "2015-03-15 19:10:25",
	Message:      "Merge pull request #2 from SimhaGitHub/readme-edits		Update README.md",
}

var ExpectedUser = st.User{
	Id:            1738270,
	Login:         "yegorch",
	Type_:         2,
	CreatedAt:     "2012-05-14 17:29:38",
	Followers:     5,
	FollowerList:  []uint64{972337, 6565138, 5877145, 5310714, 10236771},
	Following:     14,
	FollowingList: []uint64{12631, 174693, 222581, 342133, 401908, 490484, 526301, 574696, 1124763, 1263688, 1366872, 1413527, 2715158, 4090481},
	PublicGists:   1,
	PublicRepos:   4,
}

func setup[K comparable, T any](path string, ft func([]string) T, kt func(*T) K) catalog.Catalog[K, T] {
	f, _ := os.Open(path)
	defer f.Close()
	bufrd := bufio.NewReader(f)

	return catalog.New(parser.ConvertData(bufrd), ft, kt)
}

func TestMakeCatalog(t *testing.T) {
	c := setup("./commits_test.csv", st.CreateCommit, st.CommitKeyFunc)

	if !ExpectedCommit.Equal(c.View[1]) {
		t.Errorf("Expected %s, got %s", ExpectedCommit, *c.View[1])
	}
}

func TestCatalogHash(t *testing.T) {
	c := setup("./commits_test.csv", st.CreateCommit, st.CommitKeyFunc)

	current := c.Hash[st.CommitKeyFunc(&ExpectedCommit)]
	if !ExpectedCommit.Equal(current) {
		t.Errorf("Expected %s, got %s", ExpectedCommit, *current)
	}
}

func TestReadUser(t *testing.T) {
	u := setup("./users_test.csv", st.CreateUser, st.CreateUserKey)

	current := u.View[0]
	if !ExpectedUser.Equal(*current) {
		t.Errorf("Expected %s, got %s", ExpectedUser, *current)
	}
}
