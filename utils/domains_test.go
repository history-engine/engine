package utils

import (
	"github.com/fengqi/lrace"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractDomains(t *testing.T) {
	cases := map[string][]string{
		"https://mp.weixin.qq.com/s/lt5HsGMOvvcz0FoMqVIKKw":                                {"mp.weixin.qq.com", "*.weixin.qq.com", "*.qq.com"},
		"https://stackoverflow.com/questions/54070835/vuetify-creating-v-icon-dynamically": {"stackoverflow.com", "*.stackoverflow.com"},
		"https://blog.cloudflare.com/zh-cn/pingora-open-source-zh-cn":                      {"blog.cloudflare.com", "*.cloudflare.com"},
	}
	for k, want := range cases {
		give := ExtractDomains(k)
		assert.Equal(t, len(want), len(give))
		assert.True(t, lrace.ArrayCompare(want, give, false))
	}
}
