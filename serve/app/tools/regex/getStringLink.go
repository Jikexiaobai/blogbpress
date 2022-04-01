package regex

import (
	"github.com/gogf/gf/text/gregex"
)

func GetStringLink(text string) ([]string, error) {

	srcList, err := gregex.MatchAllString("\"link\":\\\"?(.*?)(\\\"|>|\\\\s+)", text)
	if err != nil {
		return nil, err
	}

	var pathList []string
	for _, i := range srcList {
		pathList = append(pathList, i[1])
	}
	return pathList, nil
}
