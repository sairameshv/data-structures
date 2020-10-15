# Static code analyzers and formatters
GOLINT := golint
GOFMT := gofmt
GOIMPORTS := goimports
GODOCDOWN := godocdown
STATIC_TOOLS := $(GOLINT) $(GOFMT) $(GOIMPORTS)
GO_FILES := `find . -name "*.go"`
FORMAT := format
DOC := doc
CLEAN := clean

$(FORMAT): $(STATIC_TOOLS)

$(GOLINT):
	@echo "Running golint on all the files present"
	@$(GOLINT) ./... |tee golint.log

$(GOFMT):
	@echo "Running gofmt on all the files present"
	@$(GOFMT) -s -w . &> gofmt.log;

$(GOIMPORTS):
	@echo "Running goimports on all the files present"
	@$(GOIMPORTS) -w $(GO_FILES) &> goimports.log;

$(DOC):
	@echo "Updating the doc"
	@$(GODOCDOWN) . > README.md

$(CLEAN):
	@rm -r *.log
