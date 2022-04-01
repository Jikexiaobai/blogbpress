package result

type Grade struct {
	GradeId     int64    `json:"id"`
	Title       string   `json:"title"`
	Icon        string   `json:"icon"`
	Integral    int      `json:"integral"`
	CreateGroup int      `json:"createGroup"`
	PostsModule []string `json:"postsModule"`
	CommonAuth  []string `json:"commonAuth"`
}
