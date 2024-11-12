package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Icon holds the schema definition for the Icon entity.
type Icon struct {
	ent.Schema
}

// Fields of the Icon.
func (Icon) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("host").MaxLen(100).Comment("域名"),
		field.String("path").MaxLen(500).Comment("完整本地文件地址"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("入库时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
	}
}

// Edges of the Icon.
func (Icon) Edges() []ent.Edge {
	return nil
}

func (Icon) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("host").Unique(),
	}
}

func (Icon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "icon"},
	}
}
