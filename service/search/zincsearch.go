package search

import (
	"context"
	"history-engine/engine/model"
	"history-engine/engine/service/zincsearch"
	"strconv"
)

var _ EngineInterface = ZincSearch{}

type ZincSearch struct{}

func NewZincSearch() ZincSearch {
	return ZincSearch{}
}

func (z ZincSearch) GetDocument(ctx context.Context, userId int64, docId string) (*model.SearchEngineDocument, error) {
	zd, err := zincsearch.GetDocument(ctx, userId, docId)
	if err != nil {
		return nil, err
	}

	sd := &model.SearchEngineDocument{
		Id: docId,
	}

	if source, ok := zd.Source.(map[string]interface{}); ok {
		if val, ok := source["url"].(string); ok {
			sd.Url = val
		}
		if val, ok := source["title"].(string); ok {
			sd.Title = val
		}
		if val, ok := source["excerpt"].(string); ok {
			sd.Excerpt = val
		}
		if val, ok := source["content"].(string); ok {
			sd.Content = val
		}
	}

	return sd, nil
}

func (z ZincSearch) PutDocument(ctx context.Context, userId int64, docId string, doc *model.SearchEngineDocument) error {
	zd := &model.ZincWriteDocument{
		Url:     doc.Url,
		Title:   doc.Title,
		Excerpt: doc.Excerpt,
		Content: doc.Content,
	}
	return zincsearch.PutDocument(ctx, userId, docId, zd)
}

func (z ZincSearch) DelDocument(ctx context.Context, userId int64, docId string) error {
	return zincsearch.DelDocument(ctx, userId, docId)
}

func (z ZincSearch) CreateIndex(ctx context.Context, userId int64) error {
	return zincsearch.CreateIndex(ctx, userId)
}

func (z ZincSearch) Search(ctx context.Context, userId int64, request model.SearchRequest) (*model.SearchResponse, error) {
	zs, err := zincsearch.EsSearch(ctx, userId, request)
	if err != nil {
		return nil, err
	}

	sr := &model.SearchResponse{
		Total: zs.Hits.Total.Value,
		Pages: make([]model.SearchResultPage, 0),
	}

	for _, item := range zs.Hits.Hits {
		version, err := strconv.Atoi(item.ID[32:])
		if err != nil {
			return nil, err
		}

		row := model.SearchResultPage{
			DocId:    item.ID,
			UniqueId: item.ID[0:32],
			Version:  version,
		}
		sr.Pages = append(sr.Pages, row)
	}

	return sr, nil
}
