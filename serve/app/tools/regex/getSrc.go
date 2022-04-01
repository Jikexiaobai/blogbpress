package regex

import "github.com/gogf/gf/text/gregex"

func GetSrcLink(text string) ([]string, error) {
	srcList, err := gregex.MatchAllString("<img.*?>", text)
	if err != nil {
		return nil, err
	}
	var pathList []string
	for _, i := range srcList {

		link, err := gregex.MatchString("src=\\\"?(.*?)(\\\"|>|\\\\s+)", i[0])
		if err != nil {
			return nil, err
		}
		pathList = append(pathList, link[1])
	}
	return pathList, nil
}
