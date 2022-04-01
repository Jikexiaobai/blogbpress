package result

type HomeDesignList struct {
	Count           int
	Key             string
	Mode            int
	Module          string
	ShowTitle       int
	Style           int
	Height          int
	BackgroundImage string
	Title           string
	CateId          int64
	RoleIds         []int
}

type WebHomeInfo struct {
	ShowTitle int         `json:"showTitle"` //
	Style     int         `json:"style"`     //
	Title     string      `json:"title"`     //
	List      interface{} `json:"list"`      //
}

//type WebHomeInfo struct {
//	Carousel []*CarouselListInfo `json:"carouselList"` //
//	Edu      []*EduListInfo      `json:"eduList"`      //
//	Video    []*VideoListInfo    `json:"videoList"`    //
//	Audio    []*AudioListInfo    `json:"audioList"`    //
//	Resource []*ResourceListInfo `json:"resourceList"` //
//	Article  []*ArticleListInfo  `json:"articleList"`  //
//	Group    []*GroupListInfo    `json:"groupList"`    //
//	Question []*QuestionListInfo `json:"questionList"` //
//	Topic    []*TopicListInfo    `json:"topicList"`    //
//	User     interface{}         `json:"userList"`     //
//	Role     []*RoleInfo         `json:"roleList"`     //
//}
