package result

type GroupJoinList struct {
	GroupId int64  ` json:"groupId"`
	Title   string `json:"title"`
}

type FileList struct {
	Total    int
	FileList []string
}
