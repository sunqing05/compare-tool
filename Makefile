.PHONY: build clean 

CUR_DIR = $(shell pwd)
BUILD_TO = $(CUR_DIR)/build
ARCH ?= amd64
OS ?= windows
0 ?= compare-tool.exe

build: clean
	@echo "project dir:" $(CUR_DIR)
	sh scripts/build.sh $(BUILD_TO) $(ARCH) $(OS) $(o)
	@echo "Build done"

clean:
	@rm -fr build/
	@echo "clean done"