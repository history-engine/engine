package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Alias holds the schema definition for the Alias entity.
type Alias struct {
	ent.Schema
}

// Fields of the Alias.
func (Alias) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id").Comment("用户id"),
		field.String("domain").MaxLen(100).Comment("域名"),
		field.String("alias").MaxLen(100).Comment("别名"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("入库时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
	}
}

// Edges of the Alias.
func (Alias) Edges() []ent.Edge {
	return nil
}

func (Alias) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "alias"},
	}
}

func (Alias) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "domain", "alias").Unique(),
	}
}
