package result

type HomeDesignList struct {
	Module          string
	ShowTitle       int
	Style           int
	Height          int
	BackgroundImage string
	List            string
	Desc            string
	Content         string
	Title           string
	VipIds          string
	EduIds          string
	AudioIds        string
	VideoIds        string
	ResourceIds     string
	ArticleIds      string
	TopicIds        string
	QuestionIds     string
}

type HFiveDesignList struct {
	Title       string
	ShowTitle   int
	Style       int
	Height      int
	List        string
	EduIds      string
	AudioIds    string
	VideoIds    string
	ResourceIds string
	ArticleIds  string
	Image       string
	Link        string
	IsPlatform  int
	TopicIds    string
	QuestionIds string
}

//title:item.title,
//showTitle:item.showTitle,
//height:item.height,
//backgroundImage:item.backgroundImage,
//list:item.list,
//eduIds:item.eduIds,
//audioIds:item.audioIds,
//videoIds:item.videoIds,
//resourceIds:item.resourceIds,
//articleIds:item.articleIds,
//topicIds:item.topicIds,
//questionIds:item.questionIds,
//style:item.style,
//image:item.image,
//link:item.link,
//isOpen:false,
//isPlatform:item.isPlatform,
