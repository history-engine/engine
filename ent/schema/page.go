package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int("user_id").Comment("用户id"),
		field.String("unique_id").MaxLen(32).Immutable().Comment("页面唯一id"),
		field.Int("version").Default(1).Comment("版本"),
		field.String("title").MaxLen(300).Comment("页面标题"),
		field.String("url").MaxLen(2048).Comment("原始地址"),
		field.String("path").MaxLen(500).Comment("完整本地文件地址"),
		field.Int("size").Default(0).Comment("文件大小"),
		field.Time("indexed_at").Default(time.Time{}).Comment("最后索引时间"),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the Page.
func (Page) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "unique_id", "version").Unique(),
	}
}

func (Page) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Page) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "page"},
	}
}
