// Code generated by ent, DO NOT EDIT.

package ent

import (
	"history-engine/engine/ent/page"
	"history-engine/engine/ent/schema"
	"history-engine/engine/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	pageMixin := schema.Page{}.Mixin()
	pageMixinFields0 := pageMixin[0].Fields()
	_ = pageMixinFields0
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescCreatedAt is the schema descriptor for created_at field.
	pageDescCreatedAt := pageMixinFields0[0].Descriptor()
	// page.DefaultCreatedAt holds the default value on creation for the created_at field.
	page.DefaultCreatedAt = pageDescCreatedAt.Default.(func() time.Time)
	// pageDescUpdatedAt is the schema descriptor for updated_at field.
	pageDescUpdatedAt := pageMixinFields0[1].Descriptor()
	// page.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	page.DefaultUpdatedAt = pageDescUpdatedAt.Default.(func() time.Time)
	// page.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	page.UpdateDefaultUpdatedAt = pageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// pageDescUniqueID is the schema descriptor for unique_id field.
	pageDescUniqueID := pageFields[2].Descriptor()
	// page.UniqueIDValidator is a validator for the "unique_id" field. It is called by the builders before save.
	page.UniqueIDValidator = pageDescUniqueID.Validators[0].(func(string) error)
	// pageDescVersion is the schema descriptor for version field.
	pageDescVersion := pageFields[3].Descriptor()
	// page.DefaultVersion holds the default value on creation for the version field.
	page.DefaultVersion = pageDescVersion.Default.(int)
	// pageDescTitle is the schema descriptor for title field.
	pageDescTitle := pageFields[4].Descriptor()
	// page.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	page.TitleValidator = pageDescTitle.Validators[0].(func(string) error)
	// pageDescURL is the schema descriptor for url field.
	pageDescURL := pageFields[5].Descriptor()
	// page.URLValidator is a validator for the "url" field. It is called by the builders before save.
	page.URLValidator = pageDescURL.Validators[0].(func(string) error)
	// pageDescPath is the schema descriptor for path field.
	pageDescPath := pageFields[6].Descriptor()
	// page.PathValidator is a validator for the "path" field. It is called by the builders before save.
	page.PathValidator = pageDescPath.Validators[0].(func(string) error)
	// pageDescSize is the schema descriptor for size field.
	pageDescSize := pageFields[7].Descriptor()
	// page.DefaultSize holds the default value on creation for the size field.
	page.DefaultSize = pageDescSize.Default.(int)
	// pageDescIndexedAt is the schema descriptor for indexed_at field.
	pageDescIndexedAt := pageFields[8].Descriptor()
	// page.DefaultIndexedAt holds the default value on creation for the indexed_at field.
	page.DefaultIndexedAt = pageDescIndexedAt.Default.(time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[3].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(int)
	// user.AdminValidator is a validator for the "admin" field. It is called by the builders before save.
	user.AdminValidator = userDescAdmin.Validators[0].(func(int) error)
}
