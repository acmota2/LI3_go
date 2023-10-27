package queries

import (
	"LI3_go/catalog"
	st "LI3_go/singleTypes"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadQueries(
	cfg *bufio.Reader,
	commits *catalog.Catalog[string, st.Commit],
	repos *catalog.Catalog[uint64, st.Repo],
	users *catalog.Catalog[uint64, st.User],
) {
	for {
		l, err := cfg.ReadString('\n')
		if err != nil {
			break
		}
		l = strings.TrimSuffix(l, "\n")
		current_line := string(l)
		query := strings.Split(current_line, " ")
		fmt.Println(enqueueQuery(query, commits, repos, users))
	}
}

func enqueueQuery(
	query []string,
	commits *catalog.Catalog[string, st.Commit],
	repos *catalog.Catalog[uint64, st.Repo],
	users *catalog.Catalog[uint64, st.User],
) string {
	var result string
	switch query[0] {
	case "1":
		result = Query1(users)
	case "2":
		result = Query2(commits, repos)
	case "3":
		result = Query3(commits, users)
	case "4":
		result = Query4(commits, users)
	case "5":
		n, _ := strconv.ParseUint(query[1], 10, 64)
		result = Query5(n, query[2]+" 00:00:00", query[3]+" 23:59:59", commits, users)
	case "6":
		n, _ := strconv.ParseUint(query[1], 10, 64)
		result = Query6(n, strings.ToLower(query[2]), commits, repos, users)
	case "7":
		result = Query7(query[1]+" 00:00:00", commits, repos)
	case "8":
		n, _ := strconv.ParseUint(query[1], 10, 64)
		result = Query8(n, query[1]+" 00:00:00", repos)
	case "9":
		n, _ := strconv.ParseUint(query[1], 10, 64)
		result = Query9(n, commits, users)
	}
	return result
}
