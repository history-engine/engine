package icon

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/net/html"
	"history-engine/engine/ent/icon"
	"history-engine/engine/library/db"
	"history-engine/engine/library/localcache"
	"history-engine/engine/library/logger"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var httpClient *http.Client

func init() {
	httpClient = http.DefaultClient
	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}

func DownloadIcon(ctx context.Context, pageUrl, pagePath string) error {
	parsed, err := url.Parse(pageUrl)
	if err != nil {
		return err
	}

	if parsed.Host == "" {
		return errors.New("host empty")
	}

	x := db.GetEngine()
	exist, err := x.Icon.Query().Where(icon.Host(parsed.Host)).Exist(ctx)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	iconPath := ""

	if newIcon, ok := favicon(parsed.Scheme, parsed.Host, pageUrl); ok {
		iconPath = newIcon
	}

	if iconPath == "" {
		if newIcon, ok := headerIcon(parsed.Host, setting.SingleFile.HtmlPath+pagePath); ok {
			iconPath = newIcon
		}
	}

	if iconPath != "" {
		_, err = x.Icon.Create().SetHost(parsed.Host).SetPath(iconPath).Save(ctx)
		localcache.GetEngine().Delete("icon:all")
		logger.Zap().Info("save host icon", zap.String("host", parsed.Host), zap.String("path", iconPath))
	}

	return err
}

// 下载默认的favicon.ico
func favicon(scheme, host, origin string) (string, bool) {
	if strings.Contains(host, "github.io") {
		host = "github.com"
	}

	iconFile := host + ".ico"
	iconPath := setting.Common.IconPath + "/" + iconFile
	if utils.FileExist(iconPath) {
		return iconFile, true
	}

	iconUrl := scheme + "://" + host + "/favicon.ico"
	req, err := http.NewRequest(http.MethodGet, iconUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Zap().Error("download icon err", zap.Error(err), zap.String("url", iconUrl))
		return "", false
	}

	if resp.StatusCode != http.StatusOK {
		return "", false
	}

	f, err := os.OpenFile(iconPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		n, err := io.Copy(f, resp.Body)
		if n != 0 && err == nil {
			return iconFile, true
		}
	}

	return "", false
}

// 从HTML文件的header标签里分析icon
// todo 没必要分析整个HTML文件
func headerIcon(host, htmlPath string) (string, bool) {
	fi, err := os.Open(htmlPath)
	if err != nil {
		logger.Zap().Error("open html file err", zap.Error(err), zap.String("file", htmlPath))
		return "", false
	}

	doc, err := html.Parse(fi)
	if err != nil {
		logger.Zap().Error("parse html file err", zap.Error(err), zap.String("file", htmlPath))
		return "", false
	}

	var faviconURL string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "link" {
			for _, attr := range n.Attr {
				if attr.Key == "rel" && (strings.Contains(attr.Val, "icon") || strings.Contains(attr.Val, "shortcut icon")) {
					for _, a := range n.Attr {
						if a.Key == "href" {
							faviconURL = a.Val
							return
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if faviconURL == "" || len(faviconURL) < 6 {
		return "", false
	}

	if faviconURL[0:5] == "data:" {
		if newPath := SaveBase64Image(host, faviconURL); newPath != "" {
			return newPath, true
		}
		return "", false
	}

	return faviconURL, true
}

// data:image/x-icon;base64,AAA...
// data:image/svg xml,<svg...</svg>
func SaveBase64Image(host, base64Data string) string {
	commaIndex := strings.Index(base64Data, ",")
	if commaIndex == -1 {
		return ""
	}

	var mimeType string
	var err error
	var binaryData []byte
	if base64Data[commaIndex-6:commaIndex] == "base64" {
		data := base64Data[commaIndex+1:]
		binaryData, err = base64.StdEncoding.DecodeString(data)
		mimeType = base64Data[5 : commaIndex-7]
		if err != nil {
			logger.Zap().Error("Failed to decode Base64 data", zap.Error(err), zap.String("base64Data", base64Data))
			return ""
		}
	} else {
		binaryData = []byte(base64Data[commaIndex+1:])
		mimeType = base64Data[5:commaIndex]
	}

	var fileExtension string
	switch mimeType {
	case "image/png":
		fileExtension = "png"
	case "image/jpeg":
		fileExtension = "jpg"
	case "image/gif":
		fileExtension = "gif"
	case "image/x-icon":
		fileExtension = "ico"
	case "image/vnd.microsoft.icon":
		fileExtension = "ico"
	case "image/svg+xml":
		fileExtension = "svg"
	case "application/octet-stream":
		if bytes.Contains(binaryData, []byte("<svg")) {
			fileExtension = "svg"
		}
	case "image/webp":
		fileExtension = "webp"
	default:
		logger.Zap().Warn("Unsupported MIME type: " + mimeType)
		return ""
	}

	// 保存文件
	filename := host + "." + fileExtension
	err = os.WriteFile(setting.Common.IconPath+"/"+filename, binaryData, 0644)
	if err != nil {
		logger.Zap().Error("write icon err", zap.Error(err))
		return ""
	}

	return filename
}
