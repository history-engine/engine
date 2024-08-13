package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSizeFormat(t *testing.T) {
	cases := map[int]string{
		111:      "111 Bytes",
		988780:   "965.61 KB",
		13248732: "12.63 MB",
		20971520: "20 MB",
	}
	for k, v := range cases {
		give := SizeFormat(k)
		assert.Equal(t, v, give)
	}
}
