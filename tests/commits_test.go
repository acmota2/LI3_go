package commits_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"LI3_go/commits"
)

func print_commit(c commits.Commit) {
	fmt.Printf(
		`repo_id: %d
		author_id: %d
		committer_id: %d
		message: %s
		commit_at: %s`,
		c.Repo_id,
		c.Author_id,
		c.Committer_id,
		c.Message,
		c.Commit_at.Format(time.UnixDate),
	)
}

func main() {
	f, err := os.Open("../misc/commits.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}
