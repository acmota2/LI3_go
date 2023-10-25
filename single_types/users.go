package single_types

import (
	"fmt"
	"reflect"
	"strconv"

	"LI3_go/csv_parser"
)

type UserType int

const (
	bot          UserType = iota
	organization UserType = iota
	user         UserType = iota
)

type User struct {
	Id            uint64
	Login         string
	Type_         UserType
	CreatedAt     string
	Followers     uint64
	FollowerList  []uint64
	Following     uint64
	FollowingList []uint64
	PublicGists   uint64
	PublicRepos   uint64
}

func CreateUserKey(u *User) uint64 {
	return u.Id
}

func CreateUser(data []string) User {
	id, _ := strconv.ParseUint(data[0], 10, 64)
	type_ := user
	switch data[2] {
	case "Bot", "bot":
		type_ = bot
	case "Organization", "organization":
		type_ = organization
	case "User", "user":
		type_ = user
	}
	followers, _ := strconv.ParseUint(data[4], 10, 64)
	follower_list := csv_parser.ParseVec(data[5])
	following, _ := strconv.ParseUint(data[6], 10, 64)
	following_list := csv_parser.ParseVec(data[7])
	public_gists, _ := strconv.ParseUint(data[8], 10, 64)
	public_repos, _ := strconv.ParseUint(data[9], 10, 64)

	return User{
		Id:            id,
		Login:         data[1],
		Type_:         type_,
		CreatedAt:     data[3],
		Followers:     followers,
		FollowerList:  follower_list,
		Following:     following,
		FollowingList: following_list,
		PublicGists:   public_gists,
		PublicRepos:   public_repos,
	}
}

func (u1 *User) Equal(u2 User) bool {
	return u1.Id == u2.Id &&
		u1.Login == u2.Login &&
		u1.Type_ == u2.Type_ &&
		u1.CreatedAt == u2.CreatedAt &&
		u1.Followers == u2.Followers &&
		reflect.DeepEqual(u1.FollowerList, u2.FollowerList) &&
		u1.Following == u2.Following &&
		reflect.DeepEqual(u1.FollowingList, u2.FollowingList) &&
		u1.PublicGists == u2.PublicGists &&
		u1.PublicRepos == u2.PublicRepos
}

func (u1 User) String() string {
	return "User{" +
		fmt.Sprintf("%d", u1.Id) + "," +
		u1.Login + "," +
		u1.CreatedAt + "," +
		fmt.Sprintf("%d", u1.Followers) + "," +
		printSlice(u1.FollowerList) + "," +
		fmt.Sprintf("%d", u1.Following) + "," +
		printSlice(u1.FollowingList) + "," +
		fmt.Sprintf("%d", u1.PublicGists) + "," +
		fmt.Sprintf("%d", u1.PublicRepos) + "}"
}

func printSlice(slice []uint64) string {
	str := "{"
	len_ := len(slice) - 1
	for i := 0; i < len_; i++ {
		str += fmt.Sprintf("%d", slice[i]) + ","
	}
	str += fmt.Sprintf("%d", slice[len_]) + "}"
	return str
}
