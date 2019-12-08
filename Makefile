
LDFLAGS += -X "main.buildVersion=$(shell git describe --tags --dirty)"
LDFLAGS += -X "main.buildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "main.buildCommit=$(shell git rev-list -1 --abbrev-commit HEAD)"
LDFLAGS += -X "main.buildBranch=$(shell git rev-parse --abbrev-ref HEAD)"


.PHONY: build
build:
	go build -ldflags '$(LDFLAGS)'

.PHONY: clean
clean:
	rm -r carl

