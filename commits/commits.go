package commits

import (
	"strconv"
	"time"
)

type Commit struct {
	Repo_id      uint64
	Author_id    uint64
	Committer_id uint64
	Commit_at    time.Time
	Message      string
}

func create_commit(data []string) Commit {
	repo_id, _ := strconv.ParseUint(data[0], 10, 64)
	author_id, _ := strconv.ParseUint(data[1], 10, 64)
	committer_id, _ := strconv.ParseUint(data[2], 10, 64)
	commit_at, _ := time.Parse("YYYY-MM-DD hh:mm:ss", data[3])

	return Commit{
		Repo_id:      repo_id,
		Author_id:    author_id,
		Committer_id: committer_id,
		Commit_at:    commit_at,
		Message:      data[4],
	}
}
