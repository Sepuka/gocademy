package structure

type User struct {
	Id      int
	GroupId int
}

type Group struct {
	Id   int
	Name string
}

type UserGroup struct {
	Id      int
	GroupId int
	Name    string
}

func join(usr []User, grp []Group) []UserGroup {
	var result = []UserGroup{}
	var groups = map[int]string{}

	for _, g := range grp {
		groups[g.Id] = g.Name
	}

	for _, u := range usr {
		result = append(result, UserGroup{Id: u.Id, Name: groups[u.GroupId], GroupId: u.GroupId})
	}

	return result
}
