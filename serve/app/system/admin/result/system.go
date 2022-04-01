package result

type SystemInfo struct {
	Base *Base `json:"base"`
	File *File `json:"file"`
}
type Base struct {
	Title          string `json:"title"`
	ChildTitle     string `json:"childTitle"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Logo           string `json:"logo"`
	Icon           string `json:"icon"`
	CurrencySymbol string `json:"currencySymbol"`
	Language       string `json:"language"`
	RecordNumber   string `json:"recordNumber"`
}
type File struct {
	FileSize  string   `json:"fileSize"`
	ImageType []string `json:"imageType"`
	AudioType []string `json:"audioType"`
	VideoType []string `json:"videoType"`
	OtherType []string `json:"otherType"`
}
