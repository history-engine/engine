package utils

import (
	"net/url"
	"strings"
)

// ExtractDomains 提取网址中的域名
// 跟前端逻辑不一样的地方是不包括跟域名
func ExtractDomains(inputURL string) []string {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return nil
	}

	domains := make(map[string]struct{})

	hostname := parsedURL.Hostname()
	domains[hostname] = struct{}{}

	parts := strings.Split(hostname, ".")
	if len(parts) >= 2 {
		rootDomain := strings.Join(parts[len(parts)-2:], ".")
		//domains[rootDomain] = struct{}{}
		domains["*."+rootDomain] = struct{}{}
	}

	if len(parts) > 2 {
		for i := 1; i < len(parts)-1; i++ {
			wildcardDomain := "*." + strings.Join(parts[i:], ".")
			domains[wildcardDomain] = struct{}{}
		}
	}

	result := make([]string, 0, len(domains))
	for domain := range domains {
		result = append(result, domain)
	}

	return result
}
