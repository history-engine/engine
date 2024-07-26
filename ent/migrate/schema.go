// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FileTypeColumns holds the columns for the "file_type" table.
	FileTypeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "suffix", Type: field.TypeString, Size: 100},
		{Name: "type", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// FileTypeTable holds the schema information for the "file_type" table.
	FileTypeTable = &schema.Table{
		Name:       "file_type",
		Columns:    FileTypeColumns,
		PrimaryKey: []*schema.Column{FileTypeColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "filetype_user_id_suffix",
				Unique:  true,
				Columns: []*schema.Column{FileTypeColumns[1], FileTypeColumns[2]},
			},
		},
	}
	// HostColumns holds the columns for the "host" table.
	HostColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "host", Type: field.TypeString, Size: 100},
		{Name: "type", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// HostTable holds the schema information for the "host" table.
	HostTable = &schema.Table{
		Name:       "host",
		Columns:    HostColumns,
		PrimaryKey: []*schema.Column{HostColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "host_user_id_host_type",
				Unique:  true,
				Columns: []*schema.Column{HostColumns[1], HostColumns[2], HostColumns[3]},
			},
		},
	}
	// PageColumns holds the columns for the "page" table.
	PageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "unique_id", Type: field.TypeString, Size: 32},
		{Name: "version", Type: field.TypeInt, Default: 1},
		{Name: "title", Type: field.TypeString, Size: 500, Default: ""},
		{Name: "url", Type: field.TypeString, Size: 2048},
		{Name: "path", Type: field.TypeString, Size: 500},
		{Name: "size", Type: field.TypeInt, Default: 0},
		{Name: "indexed_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PageTable holds the schema information for the "page" table.
	PageTable = &schema.Table{
		Name:       "page",
		Columns:    PageColumns,
		PrimaryKey: []*schema.Column{PageColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "page_user_id_unique_id_version",
				Unique:  true,
				Columns: []*schema.Column{PageColumns[1], PageColumns[2], PageColumns[3]},
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 100},
		{Name: "password", Type: field.TypeString, Size: 32},
		{Name: "admin", Type: field.TypeInt, Default: 0},
		{Name: "avatar", Type: field.TypeString, Size: 2048, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FileTypeTable,
		HostTable,
		PageTable,
		UserTable,
	}
)

func init() {
	FileTypeTable.Annotation = &entsql.Annotation{
		Table: "file_type",
	}
	HostTable.Annotation = &entsql.Annotation{
		Table: "host",
	}
	PageTable.Annotation = &entsql.Annotation{
		Table: "page",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
