
# 出力先のディレクトリ
BINDIR:=bin

# ビルド時にチェックする .go ファイル
GO_FILES:=$(shell find . -type f -name '*.go' -print)

# ビルドタスク
.PHONY: build
build: $(GO_FILES)
	@go build -o $(BINDIR)/

.PHONY: clean
clean:
	@rm -f $(BINDIR)/*
