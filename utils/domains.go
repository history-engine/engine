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

func FilterDuplicateDomains(domains []string) []string {
	var result []string
	wildcardSet := make(map[string]bool)

	// 统一转小写，收集所有通配符域名
	for _, d := range domains {
		d = strings.ToLower(strings.TrimSpace(d))
		if strings.HasPrefix(d, "*.") {
			base := d[2:] // "*.abc.com" -> "abc.com"
			wildcardSet[base] = true
		}
	}

	// 去重逻辑
	seen := make(map[string]bool)
	for _, d := range domains {
		d = strings.ToLower(strings.TrimSpace(d))
		if seen[d] {
			continue
		}
		seen[d] = true

		if strings.HasPrefix(d, "*.") {
			// 通配符域名本身保留
			result = append(result, d)
		} else {
			// 检查是否被某个通配符覆盖
			parts := strings.SplitN(d, ".", 2)
			if len(parts) == 2 {
				base := parts[1] // "www.abc.com" -> "abc.com"
				if wildcardSet[base] {
					continue // 被 "*.base" 覆盖，丢弃
				}
			}
			result = append(result, d)
		}
	}
	return result
}
