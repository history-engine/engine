package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// FileType holds the schema definition for the FileType entity.
type FileType struct {
	ent.Schema
}

// Fields of the FileType.
func (FileType) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id").Comment("用户id"),
		field.String("suffix").MaxLen(100).Comment("后缀"),
		field.Int("type").Max(2).Default(0).Comment("规则：1-include、2-exclude"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("入库时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
	}
}

// Edges of the FileType.
func (FileType) Edges() []ent.Edge {
	return nil
}

func (FileType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "filetype"},
	}
}

func (FileType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "suffix").Unique(),
	}
}
