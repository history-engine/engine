package search

import (
	"context"
	"history-engine/engine/model"
)

type EngineInterface interface {
	GetDocument(ctx context.Context, userId int64, docId string) (*model.SearchEngineDocument, error)
	PutDocument(ctx context.Context, userId int64, docId string, doc *model.SearchEngineDocument) error
	DelDocument(ctx context.Context, userId int64, docId string) error
	CreateIndex(ctx context.Context, userId int64) error
	Search(ctx context.Context, userId int64, request model.SearchRequest) (*model.SearchResponse, error)
}
