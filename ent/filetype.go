// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"history-engine/engine/ent/filetype"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// FileType is the model entity for the FileType schema.
type FileType struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// 用户id
	UserID int64 `json:"user_id,omitempty"`
	// 后缀
	Suffix string `json:"suffix,omitempty"`
	// 规则：1-include、2-exclude
	Type int `json:"type,omitempty"`
	// 入库时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 最后更新时间
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case filetype.FieldID, filetype.FieldUserID, filetype.FieldType:
			values[i] = new(sql.NullInt64)
		case filetype.FieldSuffix:
			values[i] = new(sql.NullString)
		case filetype.FieldCreatedAt, filetype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileType fields.
func (ft *FileType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case filetype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ft.ID = int64(value.Int64)
		case filetype.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ft.UserID = value.Int64
			}
		case filetype.FieldSuffix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field suffix", values[i])
			} else if value.Valid {
				ft.Suffix = value.String
			}
		case filetype.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ft.Type = int(value.Int64)
			}
		case filetype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ft.CreatedAt = value.Time
			}
		case filetype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ft.UpdatedAt = value.Time
			}
		default:
			ft.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FileType.
// This includes values selected through modifiers, order, etc.
func (ft *FileType) Value(name string) (ent.Value, error) {
	return ft.selectValues.Get(name)
}

// Update returns a builder for updating this FileType.
// Note that you need to call FileType.Unwrap() before calling this method if this FileType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FileType) Update() *FileTypeUpdateOne {
	return NewFileTypeClient(ft.config).UpdateOne(ft)
}

// Unwrap unwraps the FileType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ft *FileType) Unwrap() *FileType {
	_tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileType is not a transactional entity")
	}
	ft.config.driver = _tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FileType) String() string {
	var builder strings.Builder
	builder.WriteString("FileType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ft.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ft.UserID))
	builder.WriteString(", ")
	builder.WriteString("suffix=")
	builder.WriteString(ft.Suffix)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ft.Type))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ft.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ft.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FileTypes is a parsable slice of FileType.
type FileTypes []*FileType
