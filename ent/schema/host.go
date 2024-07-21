package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Host holds the schema definition for the Host entity.
type Host struct {
	ent.Schema
}

// Fields of the Host.
func (Host) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id").Comment("用户id"),
		field.String("host").MaxLen(100).Comment("域名"),
		field.Int("type").Max(2).Default(0).Comment("规则：1-include、2-exclude"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("入库时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
	}
}

// Edges of the Host.
func (Host) Edges() []ent.Edge {
	return nil
}

func (Host) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "host"},
	}
}

// Indexes of the Page.
func (Host) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "host", "type").Unique(),
	}
}
