package search

import (
	"context"
	"history-engine/engine/model"
)

var ElasticSearchEngine EngineInterface

type ElasticSearch struct{}

func init() {
	ElasticSearchEngine = ElasticSearch{}
}

func (e ElasticSearch) GetDocument(ctx context.Context, userId int64, docId string) (*model.SearchEngineDocument, error) {
	//TODO implement me
	panic("implement me")
}

func (e ElasticSearch) PutDocument(ctx context.Context, userId int64, docId string, doc *model.SearchEngineDocument) error {
	//TODO implement me
	panic("implement me")
}

func (e ElasticSearch) DelDocument(ctx context.Context, userId int64, docId string) error {
	//TODO implement me
	panic("implement me")
}

func (e ElasticSearch) CreateIndex(ctx context.Context, userId int64) error {
	//TODO implement me
	panic("implement me")
}

func (e ElasticSearch) Search(ctx context.Context, userId int64, request model.SearchRequest) (*model.SearchResponse, error) {
	//TODO implement me
	panic("implement me")
}
