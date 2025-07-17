# settings
IS_RELEASE ?= false

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# tools
GO = go
WINRES = $(GO) run github.com/tc-hib/go-winres@latest
UPX = upx

# output
BINARY = ./bin/$(GOOS)-$(GOARCH)

EXE_EXT = $(shell go env GOEXE)
ifeq ($(GOARCH),wasm)
	EXE_EXT = .wasm
endif
EXE = $(BINARY)/syobon-go$(EXE_EXT)

# flags
GO_GCFLAGS =
GO_LDFLAGS =
GO_FLAGS = -v

UPX_FLAGS = -f --best --lzma

ifeq ($(IS_RELEASE),true)
	GO_GCFLAGS += -dwarf=false
	GO_LDFLAGS += -s -w
	GO_FLAGS += -trimpath
	ifeq ($(GOOS),windows)
	GO_LDFLAGS += -H windowsgui
	endif
endif

GO_FLAGS += -gcflags="$(GO_GCFLAGS)" -ldflags="$(GO_LDFLAGS)" -buildvcs=true

ifneq ($(GOARCH),wasm)
	GO_FLAGS += -buildmode=pie
endif


.PHONY: all
all: clean build

.PHONY: run
run:
	$(GO) get
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) . $(GE_FLAGS)

.PHONY: run-debug
run-debug:
	$(GO) get
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run $(GO_FLAGS) . --log-level=debug $(GE_FLAGS)

.PHONY: build
build: $(BINARY)

$(BINARY):
	mkdir -p $(BINARY)

	$(GO) get
ifeq ($(GOOS),windows)
	$(WINRES) make --out ./cmd/syobon-go/rsrc
endif
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o $(EXE) .

ifeq ($(IS_RELEASE),true)
ifneq ($(GOARCH),wasm)
	$(UPX) $(UPX_FLAGS) $(EXE)
endif
endif

.PHONY: clean
clean:
	rm -rf $(BINARY)
ifeq ($(GOOS),windows)
	rm -f ./cmd/syobon-go/rsrc_windows_*.syso
endif

.PHONY: clean-all
clean-all:
	rm -rf ./bin/
ifeq ($(GOOS),windows)
	rm -f ./cmd/syobon-go/rsrc_windows_*.syso
endif
