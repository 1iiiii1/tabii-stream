APP=tabii-stream
DIST=dist

.PHONY: all clean win linux mac tidy

all: win

# Go module tidy
 tidy:
	go mod tidy

# Windows amd64 build (requires: brew install mingw-w64)
win: tidy
	@command -v x86_64-w64-mingw32-gcc >/dev/null 2>&1 || { \
	  echo "[ERR] x86_64-w64-mingw32-gcc not found. Before: brew install mingw-w64"; \
	  exit 2; }
	mkdir -p $(DIST)
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 \
	CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ \
	go build -o $(DIST)/$(APP).exe -trimpath -ldflags "-s -w" .
	@echo "[OK] Built $(DIST)/$(APP).exe"

# macOS universal (example)
mac: tidy
	mkdir -p $(DIST)
	go build -o $(DIST)/$(APP)-mac -trimpath -ldflags "-s -w" .

# Linux amd64 build (pure Go, no GUI here unless webview deps satisfied)
linux: tidy
	mkdir -p $(DIST)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(DIST)/$(APP)-linux -trimpath -ldflags "-s -w" .

clean:
	rm -rf $(DIST)
