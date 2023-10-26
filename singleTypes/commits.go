package singleTypes

import (
	"fmt"
	"strconv"
)

type Commit struct {
	Repo_id      uint64
	Author_id    uint64
	Committer_id uint64
	Commit_at    string
	Message      string
}

func CommitKeyFunc(c *Commit) string {
	return fmt.Sprintf("%d", c.Repo_id) +
		fmt.Sprintf("%d", c.Committer_id) +
		c.Commit_at +
		c.Message
}

func CreateCommit(data []string) Commit {
	repo_id, _ := strconv.ParseUint(data[0], 10, 64)
	author_id, _ := strconv.ParseUint(data[1], 10, 64)
	committer_id, _ := strconv.ParseUint(data[2], 10, 64)

	return Commit{
		Repo_id:      repo_id,
		Author_id:    author_id,
		Committer_id: committer_id,
		Commit_at:    data[3],
		Message:      data[4],
	}
}

func (c Commit) String() string {
	return "Commit{" +
		fmt.Sprintf("%d", c.Repo_id) + "," +
		fmt.Sprintf("%d", c.Author_id) + "," +
		fmt.Sprintf("%d", c.Committer_id) + "," +
		c.Commit_at + "," +
		c.Message + "}"
}

func (c1 Commit) Equal(c2 *Commit) bool {
	return c1.Author_id == c2.Author_id &&
		c1.Committer_id == c2.Committer_id &&
		c1.Commit_at == c2.Commit_at &&
		c1.Repo_id == c2.Repo_id &&
		c1.Message == c2.Message
}
