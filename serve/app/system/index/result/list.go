package result

// 标签
type TagList struct {
	TagId int64  ` json:"tagId"`
	Name  string `json:"title"`
}

type GroupJoin struct {
	GroupId int64  ` json:"id"`
	Title   string `json:"title"`
}

type FileList struct {
	Total    int
	FileList []string
}
