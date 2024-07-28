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
		field.Int64("user_id").Comment("用户id"),
		field.String("unique_id").MaxLen(32).Immutable().Comment("页面唯一id"),
		field.Int("version").Default(1).Comment("版本"),
		field.String("title").Default("").MaxLen(500).Comment("页面标题"),
		field.Text("excerpt").Default("").Comment("摘要"),
		field.Text("content").Default("").Comment("提取后的内容"),
		field.String("url").MaxLen(2048).Comment("原始地址"),
		field.String("path").MaxLen(500).Comment("完整本地文件地址"),
		field.Int("size").Default(0).Comment("文件大小"),
		field.Time("indexed_at").Default(time.Time{}).Comment("最后索引时间"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("入库时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
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

func (Page) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "page"},
	}
}
