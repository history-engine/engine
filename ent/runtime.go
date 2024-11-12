// Code generated by ent, DO NOT EDIT.

package ent

import (
	"history-engine/engine/ent/alias"
	"history-engine/engine/ent/filetype"
	"history-engine/engine/ent/host"
	"history-engine/engine/ent/icon"
	"history-engine/engine/ent/page"
	"history-engine/engine/ent/schema"
	"history-engine/engine/ent/setting"
	"history-engine/engine/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	aliasFields := schema.Alias{}.Fields()
	_ = aliasFields
	// aliasDescDomain is the schema descriptor for domain field.
	aliasDescDomain := aliasFields[2].Descriptor()
	// alias.DomainValidator is a validator for the "domain" field. It is called by the builders before save.
	alias.DomainValidator = aliasDescDomain.Validators[0].(func(string) error)
	// aliasDescAlias is the schema descriptor for alias field.
	aliasDescAlias := aliasFields[3].Descriptor()
	// alias.AliasValidator is a validator for the "alias" field. It is called by the builders before save.
	alias.AliasValidator = aliasDescAlias.Validators[0].(func(string) error)
	// aliasDescCreatedAt is the schema descriptor for created_at field.
	aliasDescCreatedAt := aliasFields[4].Descriptor()
	// alias.DefaultCreatedAt holds the default value on creation for the created_at field.
	alias.DefaultCreatedAt = aliasDescCreatedAt.Default.(func() time.Time)
	// aliasDescUpdatedAt is the schema descriptor for updated_at field.
	aliasDescUpdatedAt := aliasFields[5].Descriptor()
	// alias.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	alias.DefaultUpdatedAt = aliasDescUpdatedAt.Default.(func() time.Time)
	// alias.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	alias.UpdateDefaultUpdatedAt = aliasDescUpdatedAt.UpdateDefault.(func() time.Time)
	filetypeFields := schema.FileType{}.Fields()
	_ = filetypeFields
	// filetypeDescSuffix is the schema descriptor for suffix field.
	filetypeDescSuffix := filetypeFields[2].Descriptor()
	// filetype.SuffixValidator is a validator for the "suffix" field. It is called by the builders before save.
	filetype.SuffixValidator = filetypeDescSuffix.Validators[0].(func(string) error)
	// filetypeDescType is the schema descriptor for type field.
	filetypeDescType := filetypeFields[3].Descriptor()
	// filetype.DefaultType holds the default value on creation for the type field.
	filetype.DefaultType = filetypeDescType.Default.(int)
	// filetype.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	filetype.TypeValidator = filetypeDescType.Validators[0].(func(int) error)
	// filetypeDescCreatedAt is the schema descriptor for created_at field.
	filetypeDescCreatedAt := filetypeFields[4].Descriptor()
	// filetype.DefaultCreatedAt holds the default value on creation for the created_at field.
	filetype.DefaultCreatedAt = filetypeDescCreatedAt.Default.(func() time.Time)
	// filetypeDescUpdatedAt is the schema descriptor for updated_at field.
	filetypeDescUpdatedAt := filetypeFields[5].Descriptor()
	// filetype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	filetype.DefaultUpdatedAt = filetypeDescUpdatedAt.Default.(func() time.Time)
	// filetype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	filetype.UpdateDefaultUpdatedAt = filetypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	hostFields := schema.Host{}.Fields()
	_ = hostFields
	// hostDescHost is the schema descriptor for host field.
	hostDescHost := hostFields[2].Descriptor()
	// host.HostValidator is a validator for the "host" field. It is called by the builders before save.
	host.HostValidator = hostDescHost.Validators[0].(func(string) error)
	// hostDescType is the schema descriptor for type field.
	hostDescType := hostFields[3].Descriptor()
	// host.DefaultType holds the default value on creation for the type field.
	host.DefaultType = hostDescType.Default.(int)
	// host.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	host.TypeValidator = hostDescType.Validators[0].(func(int) error)
	// hostDescCreatedAt is the schema descriptor for created_at field.
	hostDescCreatedAt := hostFields[4].Descriptor()
	// host.DefaultCreatedAt holds the default value on creation for the created_at field.
	host.DefaultCreatedAt = hostDescCreatedAt.Default.(func() time.Time)
	// hostDescUpdatedAt is the schema descriptor for updated_at field.
	hostDescUpdatedAt := hostFields[5].Descriptor()
	// host.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	host.DefaultUpdatedAt = hostDescUpdatedAt.Default.(func() time.Time)
	// host.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	host.UpdateDefaultUpdatedAt = hostDescUpdatedAt.UpdateDefault.(func() time.Time)
	iconFields := schema.Icon{}.Fields()
	_ = iconFields
	// iconDescHost is the schema descriptor for host field.
	iconDescHost := iconFields[1].Descriptor()
	// icon.HostValidator is a validator for the "host" field. It is called by the builders before save.
	icon.HostValidator = iconDescHost.Validators[0].(func(string) error)
	// iconDescPath is the schema descriptor for path field.
	iconDescPath := iconFields[2].Descriptor()
	// icon.PathValidator is a validator for the "path" field. It is called by the builders before save.
	icon.PathValidator = iconDescPath.Validators[0].(func(string) error)
	// iconDescCreatedAt is the schema descriptor for created_at field.
	iconDescCreatedAt := iconFields[3].Descriptor()
	// icon.DefaultCreatedAt holds the default value on creation for the created_at field.
	icon.DefaultCreatedAt = iconDescCreatedAt.Default.(func() time.Time)
	// iconDescUpdatedAt is the schema descriptor for updated_at field.
	iconDescUpdatedAt := iconFields[4].Descriptor()
	// icon.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	icon.DefaultUpdatedAt = iconDescUpdatedAt.Default.(func() time.Time)
	// icon.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	icon.UpdateDefaultUpdatedAt = iconDescUpdatedAt.UpdateDefault.(func() time.Time)
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescUniqueID is the schema descriptor for unique_id field.
	pageDescUniqueID := pageFields[2].Descriptor()
	// page.UniqueIDValidator is a validator for the "unique_id" field. It is called by the builders before save.
	page.UniqueIDValidator = func() func(string) error {
		validators := pageDescUniqueID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(unique_id string) error {
			for _, fn := range fns {
				if err := fn(unique_id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// pageDescVersion is the schema descriptor for version field.
	pageDescVersion := pageFields[3].Descriptor()
	// page.DefaultVersion holds the default value on creation for the version field.
	page.DefaultVersion = pageDescVersion.Default.(int)
	// pageDescTitle is the schema descriptor for title field.
	pageDescTitle := pageFields[4].Descriptor()
	// page.DefaultTitle holds the default value on creation for the title field.
	page.DefaultTitle = pageDescTitle.Default.(string)
	// page.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	page.TitleValidator = pageDescTitle.Validators[0].(func(string) error)
	// pageDescExcerpt is the schema descriptor for excerpt field.
	pageDescExcerpt := pageFields[5].Descriptor()
	// page.DefaultExcerpt holds the default value on creation for the excerpt field.
	page.DefaultExcerpt = pageDescExcerpt.Default.(string)
	// pageDescContent is the schema descriptor for content field.
	pageDescContent := pageFields[6].Descriptor()
	// page.DefaultContent holds the default value on creation for the content field.
	page.DefaultContent = pageDescContent.Default.(string)
	// pageDescURL is the schema descriptor for url field.
	pageDescURL := pageFields[7].Descriptor()
	// page.URLValidator is a validator for the "url" field. It is called by the builders before save.
	page.URLValidator = func() func(string) error {
		validators := pageDescURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(url string) error {
			for _, fn := range fns {
				if err := fn(url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// pageDescPath is the schema descriptor for path field.
	pageDescPath := pageFields[8].Descriptor()
	// page.PathValidator is a validator for the "path" field. It is called by the builders before save.
	page.PathValidator = func() func(string) error {
		validators := pageDescPath.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_path string) error {
			for _, fn := range fns {
				if err := fn(_path); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// pageDescSize is the schema descriptor for size field.
	pageDescSize := pageFields[9].Descriptor()
	// page.DefaultSize holds the default value on creation for the size field.
	page.DefaultSize = pageDescSize.Default.(int)
	// pageDescParsedAt is the schema descriptor for parsed_at field.
	pageDescParsedAt := pageFields[10].Descriptor()
	// page.DefaultParsedAt holds the default value on creation for the parsed_at field.
	page.DefaultParsedAt = pageDescParsedAt.Default.(time.Time)
	// pageDescIndexedAt is the schema descriptor for indexed_at field.
	pageDescIndexedAt := pageFields[11].Descriptor()
	// page.DefaultIndexedAt holds the default value on creation for the indexed_at field.
	page.DefaultIndexedAt = pageDescIndexedAt.Default.(time.Time)
	// pageDescCreatedAt is the schema descriptor for created_at field.
	pageDescCreatedAt := pageFields[12].Descriptor()
	// page.DefaultCreatedAt holds the default value on creation for the created_at field.
	page.DefaultCreatedAt = pageDescCreatedAt.Default.(func() time.Time)
	// pageDescUpdatedAt is the schema descriptor for updated_at field.
	pageDescUpdatedAt := pageFields[13].Descriptor()
	// page.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	page.DefaultUpdatedAt = pageDescUpdatedAt.Default.(func() time.Time)
	// page.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	page.UpdateDefaultUpdatedAt = pageDescUpdatedAt.UpdateDefault.(func() time.Time)
	settingFields := schema.Setting{}.Fields()
	_ = settingFields
	// settingDescCreatedAt is the schema descriptor for created_at field.
	settingDescCreatedAt := settingFields[6].Descriptor()
	// setting.DefaultCreatedAt holds the default value on creation for the created_at field.
	setting.DefaultCreatedAt = settingDescCreatedAt.Default.(func() time.Time)
	// settingDescUpdatedAt is the schema descriptor for updated_at field.
	settingDescUpdatedAt := settingFields[7].Descriptor()
	// setting.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	setting.DefaultUpdatedAt = settingDescUpdatedAt.Default.(func() time.Time)
	// setting.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	setting.UpdateDefaultUpdatedAt = settingDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[4].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(int)
	// user.AdminValidator is a validator for the "admin" field. It is called by the builders before save.
	user.AdminValidator = userDescAdmin.Validators[0].(func(int) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[5].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
