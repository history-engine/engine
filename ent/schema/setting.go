package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id").Comment("用户id"),
		field.Int("max_version").Comment("最大版本数"),
		field.Int("min_version_interval").Comment("最小保存间隔（秒）"),
		field.Int("min_size").Comment("最小HTML文件大小"),
		field.Int("max_size").Comment("设置项"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("入库时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}

func (Setting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

func (Setting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "setting"},
	}
}
