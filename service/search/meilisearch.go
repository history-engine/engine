package search

import (
	"context"
	"history-engine/engine/model"
	"history-engine/engine/service/meilisearch"
	"strconv"
)

var _ EngineInterface = MeiliSearch{}

type MeiliSearch struct{}

func NewMeiliSearch() MeiliSearch {
	return MeiliSearch{}
}

// GetDocument 获取文档  todo 或许可以减少一次转换
func (m MeiliSearch) GetDocument(ctx context.Context, userId int64, docId string) (*model.SearchEngineDocument, error) {
	zd, err := meilisearch.GetDocument(ctx, userId, docId)
	if err != nil {
		return nil, err
	}

	return &model.SearchEngineDocument{
		Id:      docId,
		Url:     zd.Url,
		Title:   zd.Url,
		Excerpt: zd.Excerpt,
		Content: zd.Excerpt,
	}, err
}

// PutDocument 添加或更新文档
func (m MeiliSearch) PutDocument(ctx context.Context, userId int64, docId string, doc *model.SearchEngineDocument) error {
	md := model.MeiliDocument{
		Id:      docId,
		Url:     doc.Url,
		Title:   doc.Title,
		Excerpt: doc.Excerpt,
		Content: doc.Content,
	}
	return meilisearch.PutDocument(ctx, userId, md)
}

// DelDocument 删除文档
func (m MeiliSearch) DelDocument(ctx context.Context, userId int64, docId string) error {
	return meilisearch.DelDocument(ctx, userId, docId)
}

// CreateIndex 新建索引，可选，添加文档时会自动创建
func (m MeiliSearch) CreateIndex(ctx context.Context, userId int64) error {
	return meilisearch.CreateIndex(ctx, userId)
}

// Search 搜索，并简单处理
// todo 返回的值或许可以直接用，或者通过参数控制进返回docId
func (m MeiliSearch) Search(ctx context.Context, userId int64, request model.SearchRequest) (*model.SearchResponse, error) {
	ms, err := meilisearch.Search(ctx, userId, request)
	if err != nil {
		return nil, err
	}

	sr := &model.SearchResponse{
		Total: ms.TotalHits,
		Pages: make([]model.SearchResultPage, 0),
	}

	var version int
	for _, item := range ms.Hits {
		if len(item.Id) > 40 {
			version, err = strconv.Atoi(item.Id[40:])
		} else {
			version, err = strconv.Atoi(item.Id[32:])
		}
		if err != nil {
			return nil, err
		}
		row := model.SearchResultPage{
			DocId:    item.Id,
			UniqueId: item.Id[0:32],
			Version:  version,
		}
		sr.Pages = append(sr.Pages, row)
	}

	return sr, nil
}
