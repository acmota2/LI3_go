package single_types

import (
	"fmt"
	"strconv"
)

type Repo struct {
	Id              uint64
	OwnerId         uint64
	FullName        string
	License         string
	HasWiki         bool
	Description     string
	Language        string
	DefaultBranch   string
	CreatedAt       string
	UpdatedAt       string
	ForksCount      uint64
	OpenIssues      uint64
	StargazersCount uint64
	Size            uint64
}

func RepoKeyFunc(r *Repo) uint64 {
	return r.Id
}

func CreateRepo(data []string) Repo {
	id, _ := strconv.ParseUint(data[0], 10, 64)
	owner_id, _ := strconv.ParseUint(data[1], 10, 64)
	has_wiki := data[4] == "True" || data[3] == "true"
	forks_count, _ := strconv.ParseUint(data[10], 10, 64)
	open_issues, _ := strconv.ParseUint(data[11], 10, 64)
	stargazers_count, _ := strconv.ParseUint(data[12], 10, 64)
	size, _ := strconv.ParseUint(data[13], 10, 64)

	return Repo{
		Id:              id,
		OwnerId:         owner_id,
		FullName:        data[2],
		License:         data[3],
		HasWiki:         has_wiki,
		Description:     data[5],
		Language:        data[6],
		DefaultBranch:   data[7],
		CreatedAt:       data[8],
		UpdatedAt:       data[9],
		ForksCount:      forks_count,
		OpenIssues:      open_issues,
		StargazersCount: stargazers_count,
		Size:            size,
	}
}

func (r Repo) String() string {
	return fmt.Sprintf("%d", r.Id)
}

func (r1 Repo) Equal(r2 *Repo) bool {
	return r1.Id == r2.Id &&
		r1.OwnerId == r2.OwnerId &&
		r1.FullName == r2.FullName &&
		r1.License == r2.License &&
		r1.HasWiki == r2.HasWiki &&
		r1.Description == r2.Description &&
		r1.Language == r2.Language &&
		r1.DefaultBranch == r2.DefaultBranch &&
		r1.CreatedAt == r2.CreatedAt &&
		r1.UpdatedAt == r2.UpdatedAt &&
		r1.ForksCount == r2.ForksCount &&
		r1.OpenIssues == r2.OpenIssues &&
		r1.StargazersCount == r2.StargazersCount &&
		r1.Size == r2.Size
}
