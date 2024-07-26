NAME=history-engine
VERSION=$(shell git describe --tags || git rev-parse --short HEAD)
RELEASE_DIR=release
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-w -s -X "main.buildVersion=$(VERSION)"'

PLATFORM_LIST = \
	darwin-amd64 \
	darwin-arm64 \
	linux-amd64 \
	linux-arm64 \
	linux-arm

WINDOWS_ARCH_LIST = windows-amd64

all: linux-amd64 linux-arm64 linux-arm darwin-amd64 darwin-arm64 windows-amd64

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(RELEASE_DIR)/$(NAME)-$@
	mkdir -p $(RELEASE_DIR)/data
	cp data/zinc_template.json $(RELEASE_DIR)/data
	cp example.setting.toml $(RELEASE_DIR)/setting.toml

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(RELEASE_DIR)/$(NAME)-$@
	mkdir -p $(RELEASE_DIR)/data
	cp data/zinc_template.json $(RELEASE_DIR)/data
	cp example.setting.toml $(RELEASE_DIR)/setting.toml

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(RELEASE_DIR)/$(NAME)-$@
	mkdir -p $(RELEASE_DIR)/data
	cp data/zinc_template.json $(RELEASE_DIR)/data/
	cp example.setting.toml $(RELEASE_DIR)/setting.toml

linux-arm64:
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(RELEASE_DIR)/$(NAME)-$@
	mkdir -p $(RELEASE_DIR)/data
	cp data/zinc_template.json $(RELEASE_DIR)/data/
	cp example.setting.toml $(RELEASE_DIR)/setting.toml

linux-arm:
	GOARCH=arm GOOS=linux $(GOBUILD) -o $(RELEASE_DIR)/$(NAME)-$@
	mkdir -p $(RELEASE_DIR)/data
	cp data/zinc_template.json $(RELEASE_DIR)/data/
	cp example.setting.toml $(RELEASE_DIR)/setting.toml

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(RELEASE_DIR)/$(NAME)-$@.exe
	mkdir -p $(RELEASE_DIR)/data
	cp data/zinc_template.json $(RELEASE_DIR)/data/
	cp example.setting.toml $(RELEASE_DIR)/setting.toml

gz_releases=$(addsuffix .gz, $(PLATFORM_LIST))
zip_releases=$(addsuffix .zip, $(WINDOWS_ARCH_LIST))

$(gz_releases): %.gz : %
	chmod +x $(RELEASE_DIR)/$(NAME)-$(basename $@)
	tar -czf $(RELEASE_DIR)/$(NAME)-$(basename $@)-$(VERSION).tar.gz -C$(RELEASE_DIR) $(NAME)-$(basename $@) setting.toml data

$(zip_releases): %.zip : %
	zip -m -j $(RELEASE_DIR)/$(NAME)-$(basename $@)-$(VERSION).zip $(RELEASE_DIR)/$(NAME)-$(basename $@).exe $(RELEASE_DIR)/setting.toml $(RELEASE_DIR)/data

release: $(gz_releases) $(zip_releases)

clean:
	rm -rf $(RELEASE_DIR)/*
