package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("username").MaxLen(50).Unique().Comment("用户名"),
		field.String("email").MaxLen(100).Unique().Comment("邮箱地址"),
		field.String("password").MaxLen(32).StructTag(`json:"-"`).Comment("版本"),
		field.Int("admin").Max(1).Default(0).Comment("是否是管理员"),
		field.String("avatar").Default("").MaxLen(2048).Comment("头像"),
		field.Time("created_at").Immutable().Default(time.Now).Comment("入库时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "user"},
	}
}
