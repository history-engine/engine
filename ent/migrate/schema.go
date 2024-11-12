// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AliasColumns holds the columns for the "alias" table.
	AliasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64, Comment: "用户id"},
		{Name: "domain", Type: field.TypeString, Size: 100, Comment: "域名"},
		{Name: "alias", Type: field.TypeString, Size: 100, Comment: "别名"},
		{Name: "created_at", Type: field.TypeTime, Comment: "入库时间"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "最后更新时间"},
	}
	// AliasTable holds the schema information for the "alias" table.
	AliasTable = &schema.Table{
		Name:       "alias",
		Columns:    AliasColumns,
		PrimaryKey: []*schema.Column{AliasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "alias_user_id_domain_alias",
				Unique:  true,
				Columns: []*schema.Column{AliasColumns[1], AliasColumns[2], AliasColumns[3]},
			},
		},
	}
	// FiletypeColumns holds the columns for the "filetype" table.
	FiletypeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64, Comment: "用户id"},
		{Name: "suffix", Type: field.TypeString, Size: 100, Comment: "后缀"},
		{Name: "type", Type: field.TypeInt, Comment: "规则：1-include、2-exclude", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "入库时间"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "最后更新时间"},
	}
	// FiletypeTable holds the schema information for the "filetype" table.
	FiletypeTable = &schema.Table{
		Name:       "filetype",
		Columns:    FiletypeColumns,
		PrimaryKey: []*schema.Column{FiletypeColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "filetype_user_id_suffix",
				Unique:  true,
				Columns: []*schema.Column{FiletypeColumns[1], FiletypeColumns[2]},
			},
		},
	}
	// HostColumns holds the columns for the "host" table.
	HostColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64, Comment: "用户id"},
		{Name: "host", Type: field.TypeString, Size: 100, Comment: "域名"},
		{Name: "type", Type: field.TypeInt, Comment: "规则：1-include、2-exclude", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "入库时间"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "最后更新时间"},
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
	// IconColumns holds the columns for the "icon" table.
	IconColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "host", Type: field.TypeString, Size: 100, Comment: "域名"},
		{Name: "path", Type: field.TypeString, Size: 500, Comment: "完整本地文件地址"},
		{Name: "created_at", Type: field.TypeTime, Comment: "入库时间"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "最后更新时间"},
	}
	// IconTable holds the schema information for the "icon" table.
	IconTable = &schema.Table{
		Name:       "icon",
		Columns:    IconColumns,
		PrimaryKey: []*schema.Column{IconColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "icon_host",
				Unique:  true,
				Columns: []*schema.Column{IconColumns[1]},
			},
		},
	}
	// PageColumns holds the columns for the "page" table.
	PageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64, Comment: "用户id"},
		{Name: "unique_id", Type: field.TypeString, Size: 40, Comment: "页面唯一id"},
		{Name: "version", Type: field.TypeInt, Comment: "版本", Default: 1},
		{Name: "title", Type: field.TypeString, Size: 500, Comment: "页面标题", Default: ""},
		{Name: "excerpt", Type: field.TypeString, Size: 2147483647, Comment: "摘要", Default: ""},
		{Name: "content", Type: field.TypeString, Size: 2147483647, Comment: "提取后的内容", Default: ""},
		{Name: "url", Type: field.TypeString, Size: 2048, Comment: "原始地址"},
		{Name: "path", Type: field.TypeString, Size: 500, Comment: "完整本地文件地址"},
		{Name: "size", Type: field.TypeInt, Comment: "文件大小", Default: 0},
		{Name: "parsed_at", Type: field.TypeTime, Comment: "最后解析时间"},
		{Name: "indexed_at", Type: field.TypeTime, Comment: "最后索引时间"},
		{Name: "created_at", Type: field.TypeTime, Comment: "入库时间"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "最后更新时间"},
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
	// SettingColumns holds the columns for the "setting" table.
	SettingColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64, Comment: "用户id"},
		{Name: "max_version", Type: field.TypeInt, Comment: "最大版本数"},
		{Name: "min_version_interval", Type: field.TypeInt, Comment: "最小保存间隔（秒）"},
		{Name: "min_size", Type: field.TypeInt, Comment: "最小HTML文件大小"},
		{Name: "max_size", Type: field.TypeInt, Comment: "设置项"},
		{Name: "created_at", Type: field.TypeTime, Comment: "入库时间"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "最后更新时间"},
	}
	// SettingTable holds the schema information for the "setting" table.
	SettingTable = &schema.Table{
		Name:       "setting",
		Columns:    SettingColumns,
		PrimaryKey: []*schema.Column{SettingColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "setting_user_id",
				Unique:  false,
				Columns: []*schema.Column{SettingColumns[1]},
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 50, Comment: "用户名"},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 100, Comment: "邮箱地址"},
		{Name: "password", Type: field.TypeString, Size: 32, Comment: "版本"},
		{Name: "admin", Type: field.TypeInt, Comment: "是否是管理员", Default: 0},
		{Name: "avatar", Type: field.TypeString, Size: 2048, Comment: "头像", Default: ""},
		{Name: "created_at", Type: field.TypeTime, Comment: "入库时间"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "最后更新时间"},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AliasTable,
		FiletypeTable,
		HostTable,
		IconTable,
		PageTable,
		SettingTable,
		UserTable,
	}
)

func init() {
	AliasTable.Annotation = &entsql.Annotation{
		Table: "alias",
	}
	FiletypeTable.Annotation = &entsql.Annotation{
		Table: "filetype",
	}
	HostTable.Annotation = &entsql.Annotation{
		Table: "host",
	}
	IconTable.Annotation = &entsql.Annotation{
		Table: "icon",
	}
	PageTable.Annotation = &entsql.Annotation{
		Table: "page",
	}
	SettingTable.Annotation = &entsql.Annotation{
		Table: "setting",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
