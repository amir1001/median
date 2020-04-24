VERSION:=$(shell cat ./VERSION)
SOURCES:=./internal/median
BASE:=median

all: build test

build:
		@echo "Building $(VERSION)"
		go build -v $(SOURCES)

test:
		@echo
		@echo "Testing $(VERSION)"
		go test -v -cover $(SOURCES)

# benchmark:
# 		@echo
# 		@echo "Benchmarking $(VERSION)"
# 		go test -run=XXX -bench=. $(SOURCES)

.PHONY: build test