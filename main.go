package main

import (
	"LI3_go/catalog"
	"LI3_go/csvParser"
	"LI3_go/queries"
	st "LI3_go/singleTypes"
	"bufio"
	"log"
	"os"
	"strings"
)

func fileConfig(filePath string) *bufio.Reader {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewReader(f)
}

func filePathDebunker(cfg *bufio.Reader) map[string]string {
	filePaths := map[string]string{}
	for i := 0; i < 3; i++ {
		line, err := cfg.ReadString('\n')
		line = strings.Trim(line, "\n")

		if err != nil {
			log.Fatal(err)
		}

		switch cur := strings.Split(line, " "); cur[0] {
		case "commits":
			filePaths["commits"] = cur[1]
		case "repos":
			filePaths["repos"] = cur[1]
		case "users":
			filePaths["users"] = cur[1]
		default:
			log.Fatal("Wrong format")
		}
	}
	return filePaths
}

func main() {
	cfg := fileConfig("./LI3.cfg")
	paths := filePathDebunker(cfg)
	cFile := fileConfig(paths["commits"])
	rFile := fileConfig(paths["repos"])
	uFile := fileConfig(paths["users"])

	cRaw := csvParser.ConvertData(cFile)
	rRaw := csvParser.ConvertData(rFile)
	uRaw := csvParser.ConvertData(uFile)

	commits := catalog.New(cRaw, st.CreateCommit, st.CommitKeyFunc)
	repos := catalog.New(rRaw, st.CreateRepo, st.RepoKeyFunc)
	users := catalog.New(uRaw, st.CreateUser, st.UserKeyFunc)

	queries.ReadQueries(cfg, &commits, &repos, &users) // this will print queries
}
