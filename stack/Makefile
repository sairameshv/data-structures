# Static code analyzers and formatters
GOLINT := golint
GOFMT := gofmt
GOIMPORTS := goimports
STATIC_TOOLS := $(GOLINT) $(GOFMT) $(GOIMPORTS)
GO_FILES := `find . -name "*.go"`

format: $(STATIC_TOOLS)

$(GOLINT):
	@echo "Running golint on all the files present"
	@$(GOLINT) ./... &> golint.log

$(GOFMT):
	@echo "Running gofmt on all the files present"
	@$(GOFMT) -s -w . &> gofmt.log;

$(GOIMPORTS):
	@echo "Running goimports on all the files present"
	@$(GOIMPORTS) -w $(GO_FILES) &> goimports.log;
