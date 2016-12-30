SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
NOTEBOOKS = $(sort $(dir $(wildcard */)))

BINARY=serve

VERSION=0.0.1

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -o ${BINARY} *.go
	$(foreach dir, $(NOTEBOOKS), cp $(BINARY) $(dir);)

.PHONY: clean
clean:
	find . -name \${BINARY} -type f -delete
