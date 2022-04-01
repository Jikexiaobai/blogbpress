package es

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client

// 初始化链接es
func LoadInit() {
	// 获取配置信息
	link := g.Cfg().Get("elastic.Address")

	client, err := elastic.NewClient(elastic.SetURL(gconv.String(link)),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		g.Dump(err.Error())
	}
	EsClient = client
}

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func Insert() error {
	ctx := context.Background()
	newStudent := Student{
		Name:         "Gopher doe",
		Age:          10,
		AverageScore: 99.9,
	}

	dataJSON, err := json.Marshal(newStudent)
	js := string(dataJSON)
	_, err = EsClient.Index().Index("students").
		BodyJson(js).Do(ctx)

	if err != nil {
		g.Dump(err.Error())
	}
	return nil
}
