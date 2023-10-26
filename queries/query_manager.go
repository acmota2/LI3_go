package queries

import (
	"LI3_go/catalog"
	st "LI3_go/singleTypes"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func Query1(users *catalog.Catalog[uint64, st.User]) string {
	bot, organization, user := 0, 0, 0
	for _, u := range users.View {
		switch u.Type_ {
		case st.Bot:
			bot++
		case st.Organization:
			organization++
		case st.User_:
			user++
		}
	}
	return fmt.Sprintf(`Bot: %d
	Organization: %d
	User: %d`,
		bot,
		organization,
		user)
}

func Query2(
	commits *catalog.Catalog[string, st.Commit],
	repos *catalog.Catalog[uint64, st.Repo],
) string {
	return fmt.Sprintf("%.2f", float64(len(commits.View))/float64(len(repos.View)))
}

func Query3(
	commits *catalog.Catalog[string, st.Commit],
	users *catalog.Catalog[uint64, st.User],
) string {
	repos, count := map[uint64]bool{}, 0
	isBot := func(id uint64) bool { return users.Hash[id].Type_ == st.Bot }
	for _, c := range commits.View {
		if _, ok := repos[c.Repo_id]; isBot(c.Author_id) && !ok {
			repos[c.Repo_id] = true
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}

func Query4(
	commits *catalog.Catalog[string, st.Commit],
	users *catalog.Catalog[uint64, st.User],
) string {
	return fmt.Sprintf("%.2f", float64(len(commits.View))/float64(len(users.View)))
}

type query56Data struct {
	Id       uint64
	Login    string
	CommitQt uint64
}

func query56cmp(a query56Data, b query56Data) int {
	return cmp.Compare(a.CommitQt, b.CommitQt)
}

func mapToSlice[K comparable, T any](m map[K]T) []T {
	slice := []T{}
	for _, v := range m {
		slice = append(slice, v)
	}
	return slice
}

func Query5(
	nUsers uint64,
	startDate string,
	endDate string,
	commits *catalog.Catalog[string, st.Commit],
	users *catalog.Catalog[uint64, st.User],
) string {
	resultMap := map[uint64]query56Data{}
	for _, c := range commits.View {
		curUser := c.Author_id
		val, ok := resultMap[curUser]
		if !ok {
			resultMap[curUser] = query56Data{
				Id:       c.Author_id,
				Login:    users.Hash[curUser].Login,
				CommitQt: 0,
			}
		}
		val.CommitQt++
		resultMap[curUser] = val
	}

	querySlice := mapToSlice(resultMap)

	slices.SortStableFunc(querySlice, query56cmp)
	var result string
	for _, q := range querySlice[:nUsers] {
		result += fmt.Sprintf("%d;%s;%d\n", q.Id, q.Login, q.CommitQt)
	}

	return result
}

func Query6(
	nUsers uint64,
	lang string,
	commits *catalog.Catalog[string, st.Commit],
	repos *catalog.Catalog[uint64, st.Repo],
	users *catalog.Catalog[uint64, st.User],
) string {
	resultMap := map[uint64]query56Data{}
	for _, c := range commits.View {
		repoId := c.Repo_id
		authorId := c.Author_id
		if repo := repos.Hash[repoId]; strings.ToLower(repo.Language) == lang {
			val, ok := resultMap[authorId]
			if !ok {
				resultMap[authorId] = query56Data{
					Id:       authorId,
					Login:    users.Hash[authorId].Login,
					CommitQt: 0,
				}
			}
			val.CommitQt++
			resultMap[authorId] = val
		}
	}

	querySlice := mapToSlice(resultMap)
	slices.SortStableFunc(querySlice, query56cmp)
	var result string

	for _, q := range querySlice[:nUsers] {
		result += fmt.Sprintf("%d;%s;%d\n", q.Id, q.Login, q.CommitQt)
	}

	return result
}

func Query7(
	date string,
	commits *catalog.Catalog[string, st.Commit],
	repos *catalog.Catalog[uint64, st.Repo],
) string {
	reposResult := make(map[uint64]struct{})
	for k := range repos.Hash {
		reposResult[k] = struct{}{}
	}

	for _, c := range commits.View {
		if c.Commit_at >= date {
			delete(reposResult, c.Repo_id)
		}
	}

	var result string
	for _, r := range repos.View {
		if _, ok := reposResult[r.Id]; ok {
			result += fmt.Sprintf("%d;%s", r.Id, r.Description)
		}
	}
	return result
}

type query8Data struct {
	lang  string
	count uint64
}

func cmpQuery8(a, b query8Data) int {
	return cmp.Compare(a.count, b.count)
}

func Query8(
	nLang uint64,
	startDate string,
	repos *catalog.Catalog[uint64, st.Repo],
) string {
	resultMap := make(map[string]uint64)
	for _, r := range repos.View {
		curLang := strings.ToTitle(r.Language)
		current, ok := resultMap[curLang]
		if !ok {
			resultMap[curLang] = 0
		}
		current++
		resultMap[curLang] = current
	}

	toSort := []query8Data{}
	for k, v := range resultMap {
		toSort = append(toSort, query8Data{lang: k, count: v})
	}
	slices.SortFunc(toSort, cmpQuery8)

	var result string
	for _, q8 := range toSort[:nLang] {
		result += q8.lang
	}
	return result
}

type query9Data struct {
	user  uint64
	count uint64
}

func isFriend(id uint64, u *st.User) bool {
	for _, id2 := range u.FollowerList {
		if id == id2 {
			for _, id2 := range u.FollowingList {
				if id == id2 {
					return true
				}
			}
		}
	}
	return false
}

func cmpQuery9(a, b query9Data) int {
	return cmp.Compare(a.count, b.count)
}

func Query9(
	nUsers uint64,
	commits *catalog.Catalog[string, st.Commit],
	users *catalog.Catalog[uint64, st.User],
) string {
	resultData := map[uint64]uint64{}
	for _, c := range commits.View {
		cId := c.Committer_id
		if c.Author_id != cId && isFriend(cId, users.Hash[c.Author_id]) {
			val, ok := resultData[cId]
			if !ok {
				resultData[cId] = 0
			}
			val++
			resultData[cId] = val
		}
	}

	resultSlice := []query9Data{}
	for k, v := range resultData {
		resultSlice = append(resultSlice, query9Data{user: k, count: v})
	}
	slices.SortFunc(resultSlice, cmpQuery9)

	var result string
	for _, val := range resultSlice {
		u := users.Hash[val.user]
		login := u.Login
		result += fmt.Sprintf(
			"%d;%s;count: %d",
			val.user,
			login,
			val.count,
		)
	}
	return result
}
