package data

import (
	_ "embed"
)

//go:embed default.svg
var DefaultSvg []byte

//go:embed zinc_template.json
var ZincTemplate []byte
